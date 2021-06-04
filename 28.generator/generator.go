package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Gustavo")

	for i := 0; i < 10; i++ {
		//Precisa do <-, pois é um canal que recebe valor
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
