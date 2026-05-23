package main

func main() {
	nums := []int{1, 2, 3, 4, 5}

	//iterate over the slice using range
	for _, n := range nums {
		println(n)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	//iterate over the map using range k = key, v=value
	for k, v := range m {
		println(k, v)
	}
}


//output:
1
2
3
4
5
a 1
b 2
c 3