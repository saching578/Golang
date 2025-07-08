package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2
	arr1[3] = 3
	arr1[4] = 4

	fmt.Println(arr1)

	sum := 0
	for _, v := range arr1 {
		sum += v
	}

	fmt.Println("sum is ", sum)

	//creating and assigning value to array
	var arr2 [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr2)

	//short hand declaration
	arr3 := [5]int{100, 200, 300, 400, 500}

	fmt.Println(arr3)

	//size is evaluated by compiler at compile time
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr4)

}
