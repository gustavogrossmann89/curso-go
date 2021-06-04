package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//Isso é praticamente a herança em GO, sem passar tipo
//Fazer pessoa pessoa fica errado, pois fica um estudante com o campo pessoa dentro, com os outros campos dentro desse campo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Gustavo", "Henrique", 31, 179}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciência da Computação", "UFPR"}
	fmt.Println(e1)
	fmt.Println(e1.faculdade)
}
