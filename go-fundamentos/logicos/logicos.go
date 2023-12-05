package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50: %t, Tv32: %t, sorverte: %t, saudável: %t", tv50, tv32, sorvete, !sorvete)
	// tv50: true, Tv32: false, sorverte: true, saudável: false
}
