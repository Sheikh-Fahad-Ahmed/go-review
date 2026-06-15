package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	names := []string{"fahad", "diva", "harshita", "anant", "richie"}

	// go sayHello(names[0])
	// go sayHello(names[1])
	// go sayHello(names[2])
	// go sayHello(names[3])
	// go sayHello(names[4])

	for _, name := range names {
		go sayHello(name)
	}

	time.Sleep(1 * time.Second)

	//------------------------------------------------------------------------------------------------//
	//2. Write a function double(n int, ch chan int) that sends n * 2 into the channel.
	// In main, launch this function as a goroutine for each number in the slice [1, 2, 3, 4, 5]
	// and collect all results from the channel.
	// Print the sum of all doubled values.

	numbers := []int{1, 2, 3, 4, 5}
	doubleCh := make(chan int)

	for _, num := range numbers {
		go double(num, doubleCh)
	}

	result := []int{}
	sum := 0
	for i := range numbers {
		result = append(result, <-doubleCh)
		sum += result[i]
	}

	// for _, num := range result {
	// 	sum += num
	// }

	fmt.Printf("\n\nDouble Result: %v\nSum: %d\n\n", result, sum)

	/*------------------------------------------------------------------------------------------------

		3. Simulate downloading 5 files concurrently. Each "download" is a goroutine
		   that sleeps for a random duration between 100–500ms and then prints
		   "Downloaded: file_N.txt".After all downloads finish,
		   print "All downloads complete". The program must not
		   use time.Sleep in main to wait — use sync.WaitGroup only.

	------------------------------------------------------------------------------------------------*/
	var wg sync.WaitGroup
	waitTime := rand.N(401) + 100
	randomDuration := time.Duration(waitTime) * time.Millisecond

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go download(i, randomDuration, &wg)
	}
	wg.Wait()

	//------------------------------------------------------------------------------------------------//
	/*
		4.	Write a function slowJob(ch chan string) that sleeps for 3 seconds then sends "job done" to the channel.
			In main, launch slowJob as a goroutine and use select to either receive the result or time out after 1 second.
			Print "result: job done" if it finishes in time, or "timeout: job took too long" if it doesn't.
			Test both cases by changing the sleep duration.
			Timeout duration
			1 * time.Second
			Expected output (3s sleep > 1s timeout)
			timeout: job took too long
			Hints
			1. time.After(d) returns a channel that receives a value after duration d — use it as a select case.
			2. select picks whichever case is ready first; if the job is slow, the timeout case fires.
			3. Change slowJob sleep to 500ms to verify the success path works too.

	*/

	fmt.Printf("\n\n")
	job := make(chan string)
	go slowJob(job)

	select {
	case msg := <-job:
		fmt.Printf("result: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout: job took too long")
	}

	//------------------------------------------------------------------------------------------------//
	/*
		5.	Build a simple hit counter. Launch 100 goroutines, each calling an Increment() function that adds 1
			to a shared count variable. After all goroutines finish, print the final count.
			It must always be exactly 100. First write it without a mutex and
			run go run -race main.go to observe the race condition. Then fix it with sync.Mutex.
	*/

	var count int
	var mu sync.Mutex
	var countWg sync.WaitGroup

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		count += 1
	}

	for range 100 {
		countWg.Add(1)
		go increment(&countWg)
	}

	countWg.Wait()
	fmt.Printf("\n\nCount: %d\n", count)
}

func slowJob(ch chan string) {
	time.Sleep(3 * time.Millisecond)
	ch <- "job done"
}

func download(n int, t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(t)
	fmt.Printf("Downloaded: file_%d.txt\n", n)
}

func double(n int, ch chan<- int) {
	ch <- n * 2
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
