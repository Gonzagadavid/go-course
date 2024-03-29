package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8)) // A
	fmt.Println(notaParaConceito(6.9)) // C
	fmt.Println(notaParaConceito(2.1)) // E
}
