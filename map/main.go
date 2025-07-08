package main

import "fmt"

func main() {
	var mymap map[string]int
	if mymap == nil {
		fmt.Println("mymap is nil ")
	}
	mymap = make(map[string]int)

	fmt.Println(mymap)
	//shorthand delcsration
	mymap2 := make(map[string]int)
	fmt.Println(mymap2)

	fruits := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	fruits["orange "] = 3

	fmt.Println(fruits) //output map[apple:1 banana:2 orange :3]

	//loop through the map
	for key, val := range fruits {
		fmt.Println("key:", key, ",value:", val)
	}

	val, ok := fruits["kiwi"]

	if ok {
		fmt.Println("key is their, so value is ", val)
	} else {
		fmt.Println("key is not their")
	}

	var mymap3 map[string][]any // key is string and value is slice of type []any so we can store anything like integer string  etc

	mymap3 = make(map[string][]any)

	mymap3["banglore"] = []any{560037, 560038, 560039}

	fmt.Println(mymap3) //output map[banglore:[560037 560038 560039]]

}
