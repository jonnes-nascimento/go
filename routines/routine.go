package main

import (
	"fmt"
	"time"
)

func sayRoutine(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}

// a funcao main executa em uma go routine "principal"
func mainRoutine() {
	go sayRoutine("world") // aqui e criada uma nova go routine independente

	// caso eu comente a linha abaixo, a go routine acima nao sera executada pois
	// a go routine principal terminara antes da go routine acima ser executada (antes que chegar no fmt.Println)
	sayRoutine("hello")
}
