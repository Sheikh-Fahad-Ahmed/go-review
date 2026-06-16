package main

import "fmt"

/*

2. Write a function squares(ch chan int, n int) that
sends the square of every number from 1 to n into the channel, then closes it.
In main, launch squares as a goroutine with n = 6.
Use a for range loop to receive and print each square.
The loop must stop automatically when the channel is closed — do not hardcode the count.

Hints
1Call close(ch) after the sending loop inside squares — not in main.
2for v := range ch { } automatically stops when the channel is closed and drained.
3If you forget to close the channel, the range loop will block forever — a classic deadlock.

*/

func main() {

	n := 10
	squaresCh := make(chan int, n)

	go squares(squaresCh, n)

	i := 1
	fmt.Println("Square of ")
	for num := range squaresCh {
		fmt.Printf("%d: %d\n", i, num)
		i++
	}

	fmt.Println("Done.")

}

func squares(ch chan int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i * i
	}

	close(ch)
}
