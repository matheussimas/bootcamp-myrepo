package main

import (
	"errors"
	"fmt"
)

type conn struct {
	rd    string
	close func()
}

func redis(user string, password string) (conn, error) {
	if user == "" {
		return conn{}, errors.New("usuário vazio")
	}

	fmt.Println("abrindo conexao")

	defer func() {
		fmt.Println("dentro do redis")
	}()

	return conn{rd: "string_de_conexao", close: func() { fmt.Println("oi") }}, nil
}

func mysql() {
	fmt.Println("fechando a conexão do mysql...")
}

func mongo() {
	fmt.Println("fechando a conexão do mongo...")
}

func main() {
	connRedis, err := redis("user", "password")
	if err != nil {
		panic(err)
	}

	fmt.Println(connRedis)

	defer connRedis.close()

	defer mysql()
	defer mongo()
}
