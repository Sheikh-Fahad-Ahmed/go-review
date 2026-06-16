package main

import (
	"fmt"
	"time"
)

/*

4. Write a function ticker(done <-chan struct{}, out chan<- int) that sends an
incrementing counter (1, 2, 3…) into out every 200ms, but stops immediately when the done channel is closed.
In main, start the ticker and collect exactly 5 values from out, printing each one.
After collecting 5 values, close the done channel to stop the goroutine gracefully. Then print "ticker stopped".
Use select inside the ticker with time.After or time.Sleep for the delay.

*/

func main() {
	out := make(chan int)
	done := make(chan struct{})
	go ticker(done, out)

	for range 5 {
		fmt.Println("counter: ", <-out)
	}

	// done <- struct{}{}
	close(done)
	close(out)

	fmt.Println("ticker stopped")

}

func ticker(done <-chan struct{}, out chan<- int) {
	count := 1
	for {

		select {
		case <-time.After(200 * time.Millisecond):
			out <- count
			count++
		case <-done:
			return
		}
	}
}
