package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// 0 1 1 2 3 5 8 13 21 34 55 ...
	//Função para retornar o N-ésimo número na sequência de Fibonacci
	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
