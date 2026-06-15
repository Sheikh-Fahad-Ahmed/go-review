package order

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)



type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func GenerateOrder(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}

	return orders
}

func ProcessOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}



func ReportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- order status report ---")
		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}
		fmt.Printf("--------------------------------\n")
	}
}
