package main

import "fmt"

type Pessoa struct {
	Name string
	Age  int8
	Job  string
}

func printPessoa(pessoa Pessoa) {
	fmt.Println("Nome: ", pessoa.Name)
	fmt.Println("Idade: ", pessoa.Age)
	fmt.Println("Prof: ", pessoa.Job)
	fmt.Println("______________________")
}

func main() {
	matheus := Pessoa{
		Name: "Matheus",
		Age:  28,
		Job:  "Desenvolvedor",
	}

	var maria Pessoa

	maria.Name = "Maria"
	maria.Age = 8

	printPessoa(matheus)
	printPessoa(maria)
}
