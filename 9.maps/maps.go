package main

import "fmt"

func main() {

	//Dentro do colchetes fica o tipo da chave e fora o tipo dos valores
	//Chaves e valores são sempre do mesmo tipo
	//Não tem como um valor ser string e outro ser int
	usuario := map[string]string{
		"nome":      "Gustavo",
		"sobrenome": "Henrique",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	//Dá para aninhar maps dentro de maps
	usuario2 := map[string]map[string]string{
		"nomes": {
			"primeiro": "Gustavo",
			"segundo":  "Henrique",
		},
		"cursos": {
			"nome":      "Ciência da Computação",
			"faculdade": "UFPR",
		},
	}

	fmt.Println(usuario2)

	//Deletando uma chave dentro do map
	delete(usuario2, "nomes")
	fmt.Println(usuario2)

	//Adicionando nova chave no map
	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}

	fmt.Println(usuario2)
}
