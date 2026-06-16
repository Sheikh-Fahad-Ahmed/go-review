package main

import (
	"fmt"
	"sync"
	"time"
)

/*

3. Simulate an order processing system. You have 12 orders (order IDs 1–12) and 4 worker goroutines.
Each worker should pick up orders from a shared jobs channel and print:
"Worker W processed order O"
Each worker simulates processing time with a 50ms sleep. After all orders are processed, print "All orders processed".
Use a buffered channel for jobs, close(jobs) to stop workers, and a WaitGroup to wait for all workers.

*/

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 12)

	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go processOrder(jobs, w, &wg)
	}

	for orderID := 1; orderID <= 12; orderID++ {
		jobs <- orderID
	}

	close(jobs)

	wg.Wait()
	fmt.Println("All orders processed.")

}

func processOrder(ch <-chan int, w int, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range ch {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d processed order %d\n", w, order)

	}
}
