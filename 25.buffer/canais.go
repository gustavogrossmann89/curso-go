package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) //canal com buffer de capacidade 2 ---- se não fizer isso, vai dar deadlock nesse código
	canal <- "Olá Gustavo!"       //setando um valor no canal
	canal <- "Tudo bem?"          //setando um valor no canal

	//Recebendo as duas mensagens
	mensagem := <-canal
	mensagem2 := <-canal

	//Printando as duas
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
