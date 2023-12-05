package main

import "fmt"

// não tem operador ternario

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2)) // aprovado
	fmt.Println(obterResultado(6.0)) // aprovado
	fmt.Println(obterResultado(5.9)) // reprovado
}
