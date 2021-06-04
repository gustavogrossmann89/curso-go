package main

import "fmt"

func func1(texto string) {
	fmt.Println(texto)
}

func func2(numero int) {
	fmt.Println(numero)
}

func main() {
	defer func1("Testando")
	func2(20)
}
