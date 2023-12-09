package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 2))
}
