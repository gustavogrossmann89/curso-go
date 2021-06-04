package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return ""
	}
}

func main() {
	numero := 15

	if numero > 10 {
		fmt.Println("Numero maior que 10")
	} else {
		fmt.Println("Numero menor ou igual a 10")
	}

	//if init
	//a variável numero2 só pode ser usada dentro do IF
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero maior que 0: ", numero2)
	} else {
		fmt.Println("Numero menor ou igual a 0: ", numero2)
	}

	fmt.Println(diaDaSemana(4))
}
