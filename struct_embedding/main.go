//with embedding we can implement the inheritance in go by embedding a struct into another struct. The embedded struct's fields and methods are promoted to the outer struct, allowing us to access them directly without having to specify the embedded struct's name. This allows us to reuse code and create more complex data structures.

package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
	person    //embedding person struct into order struct
}

type vipOrder struct {
	order
	vipLevel string
}

func main() {
	//embedding is a way to reuse code in Go by embedding one struct into another struct. The embedded struct's fields and methods are promoted to the outer struct, allowing us to access them directly without having to specify the embedded struct's name.

	//we will implement a person will embedded in order struct to represent the customer who placed the order.
	newOerder := order{
		id:     "123",
		amount: 100.0,
		status: "pending",
		person: person{
			name: "Sachin Gupta",
			age:  30,
		},
	}

	//we can print the order struct and see that it has the fields of the person struct as well.

	println(newOerder.name) //Sachin Gupta
	println(newOerder.age)  //30

	fmt.Println(newOerder)

	//lets implement inheritance using embedding by creating a new struct called vipOrder that embeds the order struct and adds a new field called vipLevel.

	vipOrder1 := vipOrder{
		order: order{
			id:     "456",
			amount: 200.0,
			status: "pending",
		},
		vipLevel: "gold",
	}
	fmt.Println(vipOrder1)

}

//output
// Sachin Gupta
// 30
// {123 100 pending {0 0 <nil>} {Sachin Gupta 30}}
// {{456 200 pending {0 0 <nil>} { 0}} gold}
