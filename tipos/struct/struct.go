package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receptor(receiver)

func (p produto) precoDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoDesconto())

	produto2 := produto{"notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoDesconto())
}
