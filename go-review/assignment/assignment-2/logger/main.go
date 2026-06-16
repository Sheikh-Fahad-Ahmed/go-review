package main

import "fmt"

/*

1. You are building a simple logger. Write a function logMessage(ch chan string, msg string) that sends a message into the channel.
In main, create a buffered channel of size 5 and send these 5 log messages without launching any goroutines:
"INFO: server started", "INFO: request received", "WARN: high memory", "ERROR: db timeout", "INFO: request done"
Then read and print all messages from the channel one by one.
The key insight: no goroutine needed because the buffer absorbs all sends before any receive.

*/

func main() {
	logCh := make(chan string, 5)

	messages := []string{"INFO: server started",
		"INFO: request received",
		"WARN: high memory",
		"ERROR: db timeout",
		"INFO: request done",
	}

	for _, msg := range messages {
		logMessage(logCh, msg)
	}

	close(logCh)

	for msg := range logCh {
		fmt.Println(msg)

	}

}

func logMessage(ch chan<- string, msg string) {
	ch <- msg

}
