package main

import (
	"fmt"
	"sync"
	"time"
)

/*
5. Simulate a config loader that should only run once regardless of how many goroutines call it.
Create a struct Config with a field DSN string. Write a function loadConfig() that:
1. Prints "Loading config..."
2. Sleeps 100ms (simulating I/O)
3. Sets a package-level cfg variable to Config{DSN: "postgres://localhost/mydb"}
Wrap this in sync.Once. Launch 10 goroutines that all call GetConfig() simultaneously.
Each goroutine should print the DSN it receives. "Loading config..." must appear exactly once.

Sync Once info:
https://leangaurav.medium.com/golang-channels-vs-sync-once-for-one-time-execution-of-code-fafc81d2f54d

*/

var cfg Config

type Config struct {
	DSN string
}

func loadConfig() {
	fmt.Println("Loading config...")
	time.Sleep(100 * time.Millisecond)
	cfg = Config{
		DSN: "postgres://localhost/mydb",
	}
}

func GetConfig(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d: %s\n", id, cfg.DSN)

}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		once.Do(loadConfig)
		go GetConfig(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("done.")
}
