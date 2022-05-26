package main

import "fmt"

type Circulo struct {
	// nesse caso o raio está aberto somente para este package
	raio float32
}

// os metodos tem a funcao de atrelar funcoes a uma estrutura
// o primeiro parametro preciso referenciar qual type (struct) estou usando.
// no segundo são os paramentros de uma função comum
func (c Circulo) MetodoSetarRaio(novoRaio float32) {
	c.raio = novoRaio
	fmt.Println(c.raio)
}

// outra maneira de fazer isso é criando um "construtor"
// mas aqui eu só crio uma função global para alterar o raio, sem atrelar a struct
func NewRaio(r float32) Circulo {
	return Circulo{raio: r}
}

func main() {
	// pra poder chamar o metodo, preciso definir uma variavel com a struct
	circulo := Circulo{}

	// depois consigo chamar meu método criado, referenciando o a variavel criada acima
	circulo.MetodoSetarRaio(3.14)

	//usando a função comum para obter a troca do raio
	fmt.Println(NewRaio(10))
}
