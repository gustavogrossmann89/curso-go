package main

import (
	"avancados/formas"
	"fmt"
)

func main() {
	r := formas.Retangulo{10, 15}
	fmt.Println(r.Area())

	c := formas.Circulo{5}
	fmt.Println(c.Area())
}
