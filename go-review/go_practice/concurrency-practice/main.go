package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup

	// sayHello := func() {
	// 	defer wg.Done()
	// 	fmt.Println("hello")
	// }
	// wg.Add(1)
	// go sayHello()
	// wg.Wait()

	//--------- variables inside closures with Goroutines -----------//
	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcome"
	// }()
	// wg.Wait()
	// fmt.Println(salutation)

	//--------- variables inside closures with Goroutines with loops -----------//
	// for _, salutation := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func(salutation string) {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)
	// 	}(salutation)
	// }
	// wg.Wait()

	// -------------- sync Package: Mutex and RWmutex -----------------------------//
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrement: %d\n", count)
	}

	//increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("arithmetic complete")

}

// ------------------------------------------------------------------------------ //

// package main

// import (
// 	"fmt"
// 	"sync"

// 	"example.com/concurrency-practice/order"
// )

// func main() {

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	orderChan := make(chan *order.Order)

// 	go func(){
// 		wg.Done()
// 		for _, order := range order.GenerateOrder(20) {
// 				orderChan <- order
// 		}
// 		fmt.Println("Done with generating orders")

// 	}()

// 	go func() {
// 		defer wg.Done()
// 		order.ProcessOrders(orders)

// 	}()

// 	wg.Wait()

// 	order.ReportOrderStatus(orders)
// 	fmt.Println("All operations completed. Exiting")
// }
