package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/matheussimas/bootcamp-myrepo/cmd/server/controllers"
	"github.com/matheussimas/bootcamp-myrepo/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := controllers.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()

}
