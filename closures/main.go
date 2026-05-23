package main

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	//closure is a function that has access to the variables in its outer scope, even after the outer function has finished executing.
	increment := counter()
	println(increment()) //1
	println(increment()) //2
	println(increment()) //3
}
