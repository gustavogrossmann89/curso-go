package main

import (
	"fmt"
	"math"
)

//Essa é nossa interface
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área é: %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

//Implementação da interface, para o struct retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

//Implementação da interface, para o struct circulo
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

//O mesmo poderia ser feito para qualquer struct que vá utilizar o método area, com a mesma assinatura (mesmos parâmetros e retornos)
func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{5}
	escreverArea(c)
}
