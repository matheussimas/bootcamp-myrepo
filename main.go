package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	// eu posso ler essa struct como json assim:
	Name        string `json:"Nome"`
	Age         int8   `json:"Idade"`
	Job         string `json:"Prof"`
	Preferencia Preferencia
}

type Preferencia struct {
	Comida string `json:"Comida"`
}

func printPessoaJson(pessoaJson string, err error) {
	fmt.Println(pessoaJson, err)
}

func main() {
	matheus := Pessoa{"Matheus", 28, "Dev", Preferencia{"Macarrão"}}
	meuJson, err := json.Marshal(matheus)
	// para conseguir imprimir o json, preciso transformar em string
	meuJsonString := string(meuJson)

	printPessoaJson(meuJsonString, err)

}
