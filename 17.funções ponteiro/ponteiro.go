package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20

	//Esse primeiro caso, não altera a variável numero, apenas a imprime
	fmt.Println(numero)
	fmt.Println(inverterSinal(numero))

	//Esse segundo sim, altera
	fmt.Println(numero)
	inverterSinalPonteiro(&numero)
	fmt.Println(numero)
}
