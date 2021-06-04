package main

import "fmt"

//Demonstrando as formas de declarar variáveis
func main() {
	var variavel1 string = "Testando"
	variavel2 := "Testando 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Testando 3"
		variavel4 string = "Testando 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Testando 5", "Testando 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//Para inverter o valor das variáveis, de forma simples, sem variável auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
