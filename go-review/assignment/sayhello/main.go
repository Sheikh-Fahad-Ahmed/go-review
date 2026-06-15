package main

import (
	"fmt"
	"time"
)

/*------------------------------------------------------------------------------------------------
1.	You are given a list of 5 names. Print "Hello, <name>!" for each name concurrently
	using goroutines — do not use a for loop that calls the function directly (that would be sequential).
	The program must wait for all goroutines to finish before exiting.
	You may use time.Sleep for this first problem only.
------------------------------------------------------------------------------------------------*/

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
}
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
