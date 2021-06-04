package enderecos

import (
	"fmt"
	"testing"
)

type cenario struct {
	enderecoInserido string
	retornoEsperado  string
}

//Comandos úteis:
// go test --cover
// go test --coverprofile cobertura.txt
// go tool cover --func=cobertura.txt
// go tool cover --html=cobertura.txt

//Obrigatóriamente começa com a palavra Test, seguido de uma letra maiúscula
//Uma boa prática é ter o mesmo nome do método que se quer testar, com o Test na frente
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	//Teste de um caso específico
	/*enderecoTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoDeEndereco(enderecoTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		//t.Error("O tipo recebido é diferente do esperado!")
		t.Errorf("O tipo recebido é diferente do esperado! Esperado: %s | Recebido: %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}*/

	//Teste de vários cenários
	cenarios := []cenario{
		{"Rua de Teste", "Rua"},
		{"Avenida de Teste", "Avenida"},
		{"Rodovia de Teste", "Rodovia"},
		{"Estrada de Teste", "Estrada"},

		//2 linhas abaixo comentadas para ver o resultado do comando: go test --cover
		//O resultado deverá ser menos que 100% de cobertura da função
		//{"ABC de Teste", "Tipo inválido!"},
		//{" ", "Tipo inválido!"},
	}

	for _, cenario := range cenarios {
		tipoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoRecebido != cenario.retornoEsperado {
			//t.Error("O tipo recebido é diferente do esperado!")
			t.Errorf("O tipo recebido é diferente do esperado! Esperado: %s | Recebido: %s", cenario.retornoEsperado, tipoRecebido)
		}
	}
}

func TestSimples(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		fmt.Println("Condição impossível!")
	}
}
