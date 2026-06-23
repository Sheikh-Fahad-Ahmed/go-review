package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

type jobs struct {
	id       int
	duration time.Duration
}

type results struct {
	jobId     int
	workerId  int
	TotalJobs int64
	status    string
	err       error
}

func main() {
	var counter int64
	var wg sync.WaitGroup
	var sampleJobs []jobs
	var failedJobs int

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	for i := range 20 {
		sampleJobs = append(sampleJobs, jobs{
			id:       i + 1,
			duration: time.Duration(rand.N(5)+1) * time.Second,
		})
	}

	jobsCh := make(chan jobs, len(sampleJobs))
	resultsCh := make(chan results, len(sampleJobs))

	now := time.Now()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(ctx, i, jobsCh, resultsCh, &wg, &counter)
	}

	for _, job := range sampleJobs {
		jobsCh <- job
	}

	close(jobsCh)

	go func() {
		wg.Wait()
		close(resultsCh)

	}()

	for result := range resultsCh {
		if result.err != nil {
			failedJobs++
			fmt.Printf("error: Worker %d has failed the job ID %d: %v\n", result.workerId, result.jobId, result.err)
		}
		fmt.Printf("worker %d has %s the job ID: %d\n", result.workerId, result.status, result.jobId)
	}
	execTime := time.Since(now)

	fmt.Printf("\ntotal jobs: %d\ntotal execution time: %v\n", atomic.LoadInt64(&counter), execTime)

}

func worker(ctx context.Context, worker int, ch <-chan jobs, resultCh chan<- results, wg *sync.WaitGroup, counter *int64) {
	defer wg.Done()

	for job := range ch {
		func() {
			defer func() {
				if r := recover(); r != nil {
					resultCh <- results{
						jobId:    job.id,
						workerId: worker,
						status:   "panicked",
						err:      fmt.Errorf("error: %v", r),
					}
				}
			}()

			if job.id == 5 {
				var nilPPointer *string
				_ = *nilPPointer
			}

			if job.id == 10 {
				resultCh <- results{
					jobId:    job.id,
					workerId: worker,
					status:   "error",
					err:      errors.New("something went wrong"),
				}
			}

			select {
			case <-time.After(job.duration):
				currentCount := atomic.AddInt64(counter, 1)

				resultCh <- results{
					jobId:     job.id,
					workerId:  worker,
					TotalJobs: currentCount,
					status:    "completed",
				}
			case <-ctx.Done():
				fmt.Println("task timed out")
				return
			}
		}()
	}
}
