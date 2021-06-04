package formas

import (
	"math"
)

//Essa é nossa interface
type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

//Implementação da interface, para o struct retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

//Implementação da interface, para o struct circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
