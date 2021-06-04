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

//Esse é o nosso método, que faria o save, por exemplo
//Basta passar o parâmetro
func (u usuario) salvar() {
	fmt.Println("Salvando dados do usuário", u)
}

//Pode-se usar retorno, normal
func (u usuario) adulto() bool {
	if u.idade >= 18 {
		return true
	}
	return false
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	enderecoExemplo := endereco{"Rua 1", 100}
	u := usuario{"Nome teste", 20, enderecoExemplo}

	//Essa é a forma de chamada, apenas chamando a partir da instância da estrutura do tipo usuario
	u.salvar()
	fmt.Println("É adulto?", u.adulto())

	fmt.Println(u.idade)
	u.aniversario()
	fmt.Println(u.idade)
}
