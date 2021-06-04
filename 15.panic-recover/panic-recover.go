package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media é igual a 6!")

}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós-execução")
}
