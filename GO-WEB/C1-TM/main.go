package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

/*
Ejercicio 2:
	1. Crea dentro de la carpeta go-web un archivo llamado main.go
	2. Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
	3. Pegale al endpoint para corroborar que la respuesta sea la correcta.

Ejercicio 3:
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
	1. Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
	2. Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
	3. Genera un handler para el endpoint llamado “GetAll”.
	4. Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/

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
