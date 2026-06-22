package main

import (
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
}

func main() {
	var counter int64
	var wg sync.WaitGroup
	var sampleJobs []jobs

	for i := range 20 {
		sampleJobs = append(sampleJobs, jobs{
			id:       i + 1,
			duration: time.Duration(rand.N(5)+1) * time.Second,
		})
	}

	jobsCh := make(chan jobs, len(sampleJobs))
	resultsCh := make(chan results, len(sampleJobs))

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, jobsCh, resultsCh, &wg, &counter)
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
		fmt.Printf("worker %d has %s the job ID: %d\n", result.workerId, result.status, result.jobId)
	}

	fmt.Printf("\n total jobs: %d\n", atomic.LoadInt64(&counter))

}

func worker(worker int, ch <-chan jobs, resultCh chan<- results, wg *sync.WaitGroup, counter *int64) {
	defer wg.Done()
	for job := range ch {
		time.Sleep(job.duration)
		currentCount := atomic.AddInt64(counter, 1)
		resultCh <- results{
			jobId:     job.id,
			workerId:  worker,
			TotalJobs: currentCount,
			status:    "completed",
		}
	}
}
