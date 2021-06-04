package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	//O conteúdo após o tipo do campo, é o nome de como será convertido quando converter para JSON
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//Novo struct
	c := cachorro{"Rex", "Dálmata", 3}

	//Converter map ou struct para JSON
	jsonCachorro, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	//Isso imprimirá um array de bytes
	fmt.Println(jsonCachorro)

	//Para printar em formato legível, é necessário o pacote bytes, com o método NewBuffer
	fmt.Println(bytes.NewBuffer(jsonCachorro))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	//Converter map ou struct para JSON
	jsonCachorro2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	//Isso imprimirá um array de bytes
	fmt.Println(jsonCachorro2)

	//Para printar em formato legível, é necessário o pacote bytes, com o método NewBuffer
	fmt.Println(bytes.NewBuffer(jsonCachorro2))

	//Criando um JSON
	jsonCachorro3 := `{"nome":"Tex","raca":"Pastor","idade":7}`

	//Criando a nova instância da estrutura que vai receber o JSON convertido
	var c3 cachorro

	//Converter JSON para struct
	//O unmarshal vai converter o JSON, colocando o resultado no endereço de memória da struct
	if erro := json.Unmarshal([]byte(jsonCachorro3), &c3); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c3)
	fmt.Println(c3.Nome)

	//Criando um JSON
	jsonCachorro4 := `{"nome":"Tux","raca":"Linguiça","idade":"10"}`

	//Criando a nova instância da estrutura que vai receber o JSON convertido
	c4 := make(map[string]string)

	//Converter JSON para map
	//O unmarshal vai converter o JSON, colocando o resultado no endereço de memória do map
	if erro := json.Unmarshal([]byte(jsonCachorro4), &c4); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c4)
	fmt.Println(c4["nome"])
}
