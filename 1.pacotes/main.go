package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("gustavo.grossmann@")
	fmt.Println(erro)
}
