package main

import (
	"fmt"
	"time"
)

/*------------------------------------------------------------------------------------------------
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

------------------------------------------------------------------------------------------------*/

func main() {

	
	job := make(chan string)
	go slowJob(job)

	select {
	case msg := <-job:
		fmt.Printf("result: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout: job took too long")
	}

}

func slowJob(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "job done"
}
