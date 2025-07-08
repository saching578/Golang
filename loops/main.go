package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	k := 0
	for {
		if k == 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	//range loops
	str := "hello world"

	for i, v := range str {
		fmt.Println("index: ", i, " value: ", string(v))
	}

	count := 0

	for range str {
		fmt.Println(count)
		count++
	}
}
