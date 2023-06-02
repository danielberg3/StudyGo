package main

import (
	"fmt"
	"math"
)

func main() {
	c := circulo{10}
	escreverArea(c)

	r := retangulo{12, 13}
	escreverArea(r)
}

type formato interface {
	area() float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func escreverArea(f formato) {
	fmt.Printf("Essa é a área da figura %f\n", f.area())
}
