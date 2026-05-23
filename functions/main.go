package main

func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, string) {
	return "Go", "java", "python"
}

func main() {
	result := add(2, 3)
	println(result)

	lang1, lang2, lang3 := getLanguage()
	println(lang1, lang2, lang3)
}

//output

5
Go java python