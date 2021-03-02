package main

import (
	"fmt"
	"time"
)

func say(msg string, done chan string) {

	defer close(done)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}

	done <- "terminei"
}

// esse exemplo cria uma go routine com um channel do tipo unbuffered. isso quer dizer que, para uma go routine
// enviar uma mensagem para um channel, outra go routine deve estar esperando receber essa mensagem nesse
// mesmo channel

func mainChannel() {

	retCanal := make(chan string)

	go say("world", retCanal) // retorna uma mensagem no channel retChannel

	fmt.Println(<-retCanal) // 	aqui, a go routine principal esta esperando receber a mensagem do channel retornado pela go routine acima
}
