package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id          int64     `json:"id"`
	Nome        string    `json:"nome"`
	Sobrenome   string    `json:"sobrenome"`
	Email       string    `json:"email"`
	Idade       int       `json:"idade"`
	Altura      float32   `json:"altura"`
	Ativo       bool      `json:"ativo"`
	DataCriacao time.Time `json:"data_criacao"`
}

type Users []User

var lastID int64
var users Users

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": &users})
}

func saveUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lastID++
	user.Id = lastID

	users = append(users, user)

	file, _ := json.MarshalIndent(user, "", " ")
	_ = ioutil.WriteFile("usuarios.json", file, 0644)

	c.JSON(200, gin.H{
		"data": &users,
	})
}

func main() {

	router := gin.Default()

	router.GET("/", listUsers)
	router.POST("/", saveUser)
	router.Run()
}
