package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// inteiros
	fmt.Println(1, 2, 1000)                                 // 1 2 1000
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000)) // Literal inteiro é int

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b)) // O byte é uint8

	// com sinal... int8 int16 int32 int64

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)        // O valor máximo do int é O valor máximo do int é 9223372036854775807
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1)) // O tipo de i1 é int

	i11 := math.MaxInt32
	fmt.Println("O valor máximo do int32 é", i11)       // O valor máximo do int32 é 2147483647
	fmt.Println("O tipo de i11 é", reflect.TypeOf(i11)) // O tipo de i11 é int

	i111 := math.MaxInt16
	fmt.Println("O valor máximo do int16 é", i111)        // O valor máximo do int16 é 32767
	fmt.Println("O tipo de i111 é", reflect.TypeOf(i111)) // O tipo de i111 é int

	i1111 := math.MaxInt8
	fmt.Println("O valor máximo do int8 é", i1111)          // O valor máximo do int8 é 127
	fmt.Println("O tipo de i1111 é", reflect.TypeOf(i1111)) // O tipo de i1111 é int

	var i2 rune = 'a'                           // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2)) // O rune é int32
	fmt.Println(i2)                             // 97

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))                 // O tipo de x é float32
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // O tipo do literal 49.99 é float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo)) // O tipo de bo é bool
	fmt.Println(!bo)                                  // false

	// string
	s1 := "Olá meu nome é David"
	fmt.Println(s1 + "!")                         // Olá meu nome é David!
	fmt.Println("O tamanho da string é", len(s1)) // O tamanho da string é 22

	s2 := `Olá
	meu
	nome
	é
	David
	`
	fmt.Println(s2 + "!")
	/*
		Olá
		meu
		nome
		é
		David
		!
	*/
	fmt.Println("O tamanho da string é", len(s2)) // O tamanho da string é 28

	// char ??
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) // O tipo de char é int32
	fmt.Println(char)                                     // 97
}
