package main

import "fmt"

func sum(s []int, c chan int) {

	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// cria duas go routines para somar o conteudo do slice, cada uma somando metade do slice
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // recebe de c (neste caso, os valores de c foram empilhados e funcionam no padrao lifo)

	fmt.Println(x, y, x+y)
}
