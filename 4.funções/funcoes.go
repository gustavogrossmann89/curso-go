package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Função com retorno múltiplo
func calculos(n1, n2 int8) (int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	return soma, subtracao, multiplicacao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto passado por parâmetro")
	fmt.Println(resultado)

	//Função com retorno múltiplo
	//Caso não quisessemos utilizar algum retorno, basta substituir a variável de recebimento por underline _
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao := calculos(5, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao)
}
