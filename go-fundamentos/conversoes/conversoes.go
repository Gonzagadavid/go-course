package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // 1.2
	fmt.Println(int(x) / y)     // 1

	nota := 6.9

	notaFinal := int(nota)

	fmt.Println(notaFinal) // 6

	// cuidado...
	fmt.Println("Teste " + string(97)) // Teste a

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // Teste 123

	// string para int
	num, err := strconv.Atoi("123")
	fmt.Println(num - 122) // 1
	fmt.Println(err)       // <nil>

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
	// Verdadeiro
}
