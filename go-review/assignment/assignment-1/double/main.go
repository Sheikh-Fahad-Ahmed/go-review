package main

import "fmt"

	/*------------------------------------------------------------------------------------------------

		2.	Write a function double(n int, ch chan int) that sends n * 2 into the channel.
	 		In main, launch this function as a goroutine for each number in the slice [1, 2, 3, 4, 5]
	 		and collect all results from the channel.
	 		Print the sum of all doubled values.

	------------------------------------------------------------------------------------------------*/

func main() {
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
}

func double(n int, ch chan<- int) {
	ch <- n * 2
}
