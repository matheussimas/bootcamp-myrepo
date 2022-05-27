package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id          int       `json:"id"`
	Nome        string    `json:"nome"`
	Sobrenome   string    `json:"sobrenome"`
	Email       string    `json:"email"`
	Idade       int       `json:"idade"`
	Altura      float32   `json:"altura"`
	Ativo       bool      `json:"ativo"`
	DataCriacao time.Time `json:"data_criacao"`
}

func writeJsonFile(c *gin.Context) {
	file, _ := json.MarshalIndent(c, "", " ")
	err := ioutil.WriteFile("/go-web/user.json", file, 0644)

	if err != nil {
		log.Fatal(err)
	}

}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deuboa",
	})
}

func main() {

	router := gin.Default()

	router.GET("/", hello)
	router.POST("/", writeJsonFile)
	router.Run()
}
