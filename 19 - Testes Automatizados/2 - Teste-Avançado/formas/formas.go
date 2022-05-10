package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	raio float64
}

func (c circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

}