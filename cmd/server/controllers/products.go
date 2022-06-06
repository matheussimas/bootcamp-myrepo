package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheussimas/bootcamp-myrepo/internal/products"
)

type ProductController struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *ProductController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Store(req.Name, req.Typee, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

type request struct {
	Name  string  `json:"name"`
	Typee string  `json:"typee"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}
