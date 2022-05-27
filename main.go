package main

import (
	"fmt"
	"math"
)

// a interface tem a intencao de criar funcoes que serao 'compartilhadas'
// para usos similares mas nao exatamente iguais
// ex: quadrado e circulo possuem areas e perimetros, porem sao calculados de formas diferentes
type Geometria interface {
	area() float32
	perimetro() float32
}

type Circulo struct {
	raio float32
}
type Retangulo struct {
	base   float32
	altura float32
}

// o parametro de receiver que vai definir quando essa funcao sera chamada
// ex: toda vez que eu passar uma struct de Circulo para a funcao area, vou acionar a area do circulo,
// pois sao funcoes diferentes, só compartilham o mesmo nome
func (c Circulo) area() float32 {
	return math.Pi * c.raio * c.raio
}

func (c Circulo) perimetro() float32 {
	return 2 * math.Pi * c.raio
}

// mesma coisa acontece com as funcs de retangulo
func (r Retangulo) perimetro() float32 {
	return r.base*2 + r.altura*2
}

func (r Retangulo) area() float32 {
	return r.base * r.altura
}

// aqui eu aplico uma funcao que vai receber a interface
// o go vai chamar cada funcao dependendo do parametro passado
func detalhe(g Geometria) {
	fmt.Println(g)
	println(g.area())
	println(g.perimetro())
}

func main() {
	c := Circulo{raio: 12}
	r := Retangulo{base: 10, altura: 3}

	// se eu passar o c (struct Circulo) como parametro ele vai acionar as funcoes de geometria do Circulo
	detalhe(c)
	// se eu passar o r (struct Retangulo) ele vai acionar as funcoes de geometria do Retangulo
	detalhe(r)
}
