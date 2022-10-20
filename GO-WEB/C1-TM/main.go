package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ejercicio 2:
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Juan",
		})
	})
	router.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Juan",
		})
	})
	// Ejercicio 3:
	type Products struct {
		ID            int     `json:"id"`
		Nombre        string  `json:"nombre"`
		Color         string  `json:"color"`
		Precio        float64 `json:"precio"`
		Stock         bool    `json:"stock"`
		Codigo        string  `json:"codigo"`
		Publicado     string  `json:"publicado"`
		FechaCreación string  `json:"fecha_creacion"`
	}
	router.GET("/products", func(ctx *gin.Context) {
		file, err := os.ReadFile("products.json")
		var products []Products
		if err != nil {
			panic("Se presentó un error al leer el archivo")
		}
		if err := json.Unmarshal(file, &products); err != nil {
			panic("Se presentó un error al decodificar el archivo")
		}
		ctx.JSON(200, gin.H{
			"products": products,
		})
	})
	router.Run()
}
