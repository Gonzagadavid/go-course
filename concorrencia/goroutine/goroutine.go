package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3) // executa sequencialmente
	// fale("João", "Só posso falar depois de você!", 1)

	// executa de forma concorrente, porém enquanto a thread principal está em execução
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	// time.Sleep(time.Second * 5) // o código acima só é executado durante esse tempo
	// fmt.Println("Fim!")

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
