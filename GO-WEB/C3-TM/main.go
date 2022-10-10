package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	ProductType string  `json:"product_type" binding:"required"`
	Count       int     `json:"count" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, productType string, count int, price float64) (Product, error)
	LastId() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
}

type Service interface {
	GetAll() ([]Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
}

func (c *Repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, ProductType: productType, Count: count, Price: price}
	updated := false
	for i := range p {
		
	}
	return p, nil
}
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func main() {
	server := gin.Default()
	server.Run()
}
