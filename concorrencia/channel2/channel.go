package main

import (
	"fmt"
	"time"
)

// channel - é a forma de comunicação entre as goroutines

func doisTresQuatro(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go doisTresQuatro(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")

	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println("C")

}
