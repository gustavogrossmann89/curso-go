package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Nome teste"
	u.idade = 18
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua 1", 100}

	u2 := usuario{"Nome teste 2", 20, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{idade: 30}
	fmt.Println(u3)
}
