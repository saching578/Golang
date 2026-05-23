package main

func add(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
func main() {
	//implementing variadic function
	sum := add(1, 2, 3, 4, 5)
	println(sum)

	//in slices
	numbers := []int{1, 2, 3, 4, 5}
	sum2 := add(numbers...)
	println(sum2)

}
