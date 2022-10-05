package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type entidad struct {
	Id       int     `json:"id" binding:"required"`
	Nombre   string  `json:"nombre" binding:"required"`
	Tipo     string  `json:"tipo" binding:"required"`
	Cantidad int     `json:"cantidad" binding:"required"`
	Precio   float64 `json:"precio" binding:"required"`
}

var (
	entidades []entidad
	lastId    int
)

func main() {
	r := gin.Default()
	r.GET("/ping", Ping())

	pr := r.Group("/products")
	pr.POST("", Guardar())
	r.Run()
}
func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "No tiene permisos para realizar la petición solicitada",
			})
			return
		}
		var req entidad
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Nombre == "" {
				ctx.JSON(400, gin.H{"error": "El campo nombre es requerido"})
				return
			}
			if req.Tipo == "" {
				ctx.JSON(400, gin.H{"error": "El campo tipo es requerido"})
				return
			}
			if req.Cantidad == 0 {
				ctx.JSON(400, gin.H{"error": "El campo cantidad es requerido"})
				return
			}
			if req.Precio == 0.0 {
				ctx.JSON(400, gin.H{"error": "El campo precio es requerido"})
				return
			}
		}
		lastId++
		req.Id = lastId
		entidades = append(entidades, req)
		ctx.JSON(200, req)
	}
}
