package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Gustavo", "Cassio", "Victor"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//Forma bem comum de se fazer
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		//Imprime os códigos da tabela ASCII
		fmt.Println(indice, letra)
	}

	for indice, letra := range "PALAVRA" {
		//Imprime a respectiva letra do código da tabela ASCII
		fmt.Println(indice, string(letra))
	}

	//FOR com MAP
	usuarios := map[string]string{
		"nome1": "Gustavo",
		"nome2": "Cassio",
	}

	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
	}

	//Não tem como fazer for com STRUCT
}
