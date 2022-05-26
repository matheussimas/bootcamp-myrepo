package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (c *Circulo) Area() float64 {
	return math.Pi * c.raio * c.raio

}
func (c *Circulo) Diametro() float64 {
	return 2 * math.Pi * c.raio
}

func main() {
	valorRaio := 5
	calculaArea := Circulo{raio: float64(valorRaio)}

	fmt.Println(calculaArea.Area())
}
