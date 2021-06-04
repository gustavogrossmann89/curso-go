package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	go escrever("Olá Gustavo") //goroutine
	// sem o go acima, o método abaixo nunca vai ser chamado
	// o go, na prática, faz com que o código abaixo não espere a chamada acima terminar, para começar a ser executado
	escrever("Tudo bem?")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
