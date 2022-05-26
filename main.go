package main

import (
	"fmt"
	"reflect"
)

type Pessoa struct {
	// eu posso ler essa struct como json assim:
	Name        string `json:"nome"`
	Age         int8   `json:"idade"`
	Job         string `json:"prof"`
	Preferencia Preferencia
}

type Preferencia struct {
	Comida string `json:"Comida"`
}

func main() {
	pessoa := Pessoa{}
	// o reflect traduz o tipo, o nome da s
	p := reflect.TypeOf(pessoa)

	for i := 0; i < p.NumField(); i++ {
		// o field traz o nome do campo do json
		field := p.Field(i)
		// tag traz o informações daqule campo na struct
		tag := field.Tag.Get("json")
		fmt.Println(tag)
		fmt.Println(field)
	}
}
