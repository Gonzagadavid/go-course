package main

import "fmt"

func closure() (funcao func()) {
	x := 10
	funcao = func() {
		fmt.Println(x)
	}

	return
}

func main() {
	x := 20
	fmt.Println(x)

	imprimex := closure()
	imprimex()
}
