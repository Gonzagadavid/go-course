package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1546.80,
			"Guga Pereira":   8457.45,
		},
		"J": {
			"João José": 11355.58,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	/*
		G map[Gabriela Silva:1546.8 Guga Pereira:8457.45]
		J map[João José:11355.58]
	*/
}
