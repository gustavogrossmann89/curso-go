package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Gustavo"), escrever("Tudo bom?"))

	for i := 0; i < 10; i++ {
		//Precisa do <-, pois é um canal que recebe valor
		fmt.Println(<-canal)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //Esse rand é só pra ser num tempo aleatório, para teste
		}
	}()

	return canal
}
