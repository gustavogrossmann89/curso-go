package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //2 goroutines

	go func() {
		escrever("Olá Gustavo")
		waitGroup.Done() // -1, vai diminuir da qtde de goroutines declaradas acima
	}()

	go func() {
		escrever("Tudo bem?")
		waitGroup.Done() // -1, vai diminuir da qtde de goroutines declaradas acima
	}()

	// Isso obriga que as duas terminem
	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
