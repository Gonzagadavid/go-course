package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i /// pegando o endereço da variavel

	*p++ // desrefenciando(pegando o valor)
	i++

	// Go não tem aritimetica de ponteiros

	fmt.Println(p, *p, i, &i) // 0xc0000a6000 3 3 0xc0000a6000

}
