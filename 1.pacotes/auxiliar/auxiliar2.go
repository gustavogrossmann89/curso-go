package auxiliar

import "fmt"

//Função privada que só pode ser chamada dentro do mesmo package auxiliar, não no main
func escrever2() {
	fmt.Println("Arquivo auxiliar2")
}
