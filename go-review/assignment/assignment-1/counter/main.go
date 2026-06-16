package main

import (
	"fmt"
	"sync"
)

/*------------------------------------------------------------------------------------------------
5.	Build a simple hit counter. Launch 100 goroutines, each calling an Increment() function that adds 1
	to a shared count variable. After all goroutines finish, print the final count.
	It must always be exactly 100. First write it without a mutex and
	run go run -race main.go to observe the race condition. Then fix it with sync.Mutex.
------------------------------------------------------------------------------------------------*/

func main() {

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
	fmt.Printf("Count: %d\n", count)
}
