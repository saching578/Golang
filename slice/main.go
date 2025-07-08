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

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice4 := slice3 // reference copy

	slice4[0] = 2

	fmt.Println(slice3) // output [2 2 3 4 5 6 7 8 9]

	slice5 := slice3[:5] // slice[0:5]

	slice5[0] = 100

	slice5 = append(slice5, 600)

	fmt.Println(slice5) //output [100 2 3 4 5 600]

}
