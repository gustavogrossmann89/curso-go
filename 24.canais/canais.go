package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //indica que só vai poder trafegar dados do tipo string no canal

	go escrever("Olá Gustavo!", canal)

	//Forma mais manual
	/*for {
		mensagem, aberto := <-canal //canal recebendo um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/

	//Forma mais automatizada/simples
	//Enquanto o canal estiver aberto e recebendo mensagens
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //setando um valor no canal
		time.Sleep(time.Second)
	}

	close(canal)
}
