package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados

	aprovados := make(map[int]string) // map[12345678978:Maria 45615487545:Pedro 78758454145:Ana]

	aprovados[12345678978] = "Maria"
	aprovados[45615487545] = "Pedro"
	aprovados[78758454145] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	/*
	   Maria (CPF: 12345678978)
	   Pedro (CPF: 45615487545)
	   Ana (CPF: 78758454145)
	*/

	fmt.Println(aprovados[78758454145]) // Ana
	delete(aprovados, 78758454145)
	fmt.Println(aprovados) // map[12345678978:Maria 45615487545:Pedro]
}
