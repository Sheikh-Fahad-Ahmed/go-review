package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*------------------------------------------------------------------------------------------------

	3. Simulate downloading 5 files concurrently. Each "download" is a goroutine
	   that sleeps for a random duration between 100–500ms and then prints
	   "Downloaded: file_N.txt".After all downloads finish,
	   print "All downloads complete". The program must not
	   use time.Sleep in main to wait — use sync.WaitGroup only.

------------------------------------------------------------------------------------------------*/

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go download(i, &wg)
	}
	wg.Wait()
	fmt.Printf("\nAll downloads complete\n")

}

func download(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	waitTime := rand.N(401) + 100
	randomDuration := time.Duration(waitTime) * time.Millisecond
	time.Sleep(randomDuration)
	fmt.Printf("Downloaded: file_%d.txt\n", n)
}
