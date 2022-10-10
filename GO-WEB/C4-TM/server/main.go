package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jcsalgadomeza/backpack-bcgow6-camilo-salgado/GO-WEB/C4-TM/products"
	"github.com/jcsalgadomeza/backpack-bcgow6-camilo-salgado/GO-WEB/C4-TM/server/handler"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	server := gin.Default()
	pr := server.Group("produts")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	if err := server.Run(); err != nil {
		panic(err)
	}
}
