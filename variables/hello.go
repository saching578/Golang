package main

import (
	"fmt"
)

// variable declaration
func main() {
	var name string = "sachin" // can be used outside function as well
	var age = 25
	dob := "04/04/1999" // cannot used outside function

	fmt.Println("Name is ", name)
	fmt.Println("age is ", age)
	fmt.Println("dob is ", dob)

	var a, b, c, d int = 1, 3, 5, 7 // multiple declaration of same type

	var Name, Age = "sachin", 27 // multiple declaration of different types

	dob, place, mobileNo := "04/04/1999", "mau", 7007926017 // multiple declaration of different types

	var (
		fatherName = "Rajendra prasad Gupta"
		motherName = "Saraswati Gupta"
	)

	fmt.Println(a, b, c, d)
	fmt.Println(Name, Age)
	fmt.Println(dob, place, mobileNo)
	fmt.Println(fatherName, motherName)
}
