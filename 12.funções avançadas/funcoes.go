package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

//Qtde de parâmetros dinâmico
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(calculos(1, 2))
	fmt.Println(soma(1, 2, 3, 4, 5))
	escrever("Texto", 1, 2, 3, 4, 5, 6)

	//Função anônima 1
	func(texto string) {
		fmt.Println(texto)
	}("Olá")

	//Função anônima 2
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebeu a string: %s", texto)
	}("Olá")

	fmt.Println(retorno)
}
