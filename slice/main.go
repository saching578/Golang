package main

import (
	"fmt"
	"math/rand"
)

// slice is a expandable array

// make--> to create slice
func main() {

	var slice1 []int //declare a slice

	fmt.Println(slice1)

	if slice1 == nil {
		fmt.Println("slice is nil because it is not initialised")
	}

	slice1 = make([]int, 5) // initialze a slice

	fmt.Println(slice1)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println(slice1)

	//shorthand delcaration
	slice2 := []int{10, 20, 30, 40, 50} // slice with values
	fmt.Println(slice2)

	// to append element at the end of a slice
	slice2 = append(slice2, 60)
	slice2 = append(slice2, 70)
	slice2 = append(slice2, 80)

	fmt.Println(slice2)

}
