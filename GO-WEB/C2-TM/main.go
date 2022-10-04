package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Ejercicio 1:
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
 1. Crea un endpoint mediante POST el cual reciba la entidad.
 2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
    deberán ir guardando todas las peticiones que se vayan realizando.
 3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
    buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
    nuevo registro (sin tener una variable de último ID a nivel global).

Ejercicio 2:
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
 1. Se debe validar todos los campos enviados en la petición, todos los campos son
    requeridos.
 2. En caso que algún campo no esté completo se debe retornar un código de error 400
    con el mensaje “el campo %s es requerido”.
    (En %s debe ir el nombre del campo que no está completo).

Ejercicio 3:
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::
 1. Al momento de enviar la petición se debe validar que un token sea enviado
 2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
 3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
    mensaje que “no tiene permisos para realizar la petición solicitada”.
*/
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
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastId++
		req.Id = lastId
		entidades = append(entidades, req)
		ctx.JSON(200, req)
	}
}
