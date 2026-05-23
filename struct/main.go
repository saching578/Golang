package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
}

// receive type
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

func (o order) getAmount() float64 {
	return o.amount
}

func newOrder(id string, amount float64) *order {
	//intial setup goes here
	myOrder := order{
		id:     id,
		amount: amount,
		status: "pending",
	}
	return &myOrder
}

func main() {
	//we don't have classes in go so to we uses structs to create objects and methods to define the behavior of those objects.
	myOrder := order{
		id:     "123",
		amount: 100.0,
		status: "pending",
	}

	//we can add later the createdAt field to the order struct and set it to the current time when we create a new order.
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)

	//implementing changeStatus method for the order struct to change the status of the order.
	myOrder.changeStatus("shipped")
	fmt.Println(myOrder)

	//implement the getAmount method which don't use pointer receiver because we don't need to modify the order struct, we just want to return the amount of the order.
	fmt.Println(myOrder.getAmount())

	//implement constructor using a function that returns a new instance of the order struct.
	newOrder := newOrder("456", 200.0)
	fmt.Println(newOrder)

	//shorthand(inlined) for using constructors

	order1 := struct {
		id     string
		amount float64
	}{
		id:     "789",
		amount: 300.0,
	}
	fmt.Println(order1)
}
