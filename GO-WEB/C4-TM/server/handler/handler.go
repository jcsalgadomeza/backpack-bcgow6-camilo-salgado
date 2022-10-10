package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcsalgadomeza/backpack-bcgow6-camilo-salgado/GO-WEB/C4-TM/products"
)

type ProductRequest struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre" binding:"required"`
	Tipo     string  `json:"tipo" binding:"required"`
	Cantidad int     `json:"cantidad" binding:"required"`
	Precio   float64 `json:"precio" binding:"required"`
}

type ProductRequestPatch struct {
	Nombre string `json:"nombre" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		p, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(p) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "No hay productos registrados",
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		var req ProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Campo nombre es requerido",
			})
			return
		}
		if req.Tipo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Campo tipo es requerido",
			})
			return
		}
		if req.Cantidad == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Campo cantidad es requerido",
			})
			return
		}
		if req.Precio <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Campo precio es requerido",
			})
			return
		}
	}
}
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Id invalido - " + err.Error(),
			})
			return
		}
		var req ProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El campo nombre es requerido",
			})
			return
		}
		if req.Tipo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El campo tipo es requerido",
			})
			return
		}
		if req.Cantidad == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El campo cantidad es requerido",
			})
			return
		}
		if req.Precio <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El campo precio es requerido",
			})
			return
		}
		p, err := p.service.Update(id, req.Nombre, req.Tipo, req.Cantidad, req.Precio)
		ctx.JSON(http.StatusOK, p)
	}
}
func (p *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Id invalido - " + err.Error(),
			})
			return
		}
		var req ProductRequestPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := p.service.UpdateName(id, req.Nombre)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)

	}
}
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id invalido - " + err.Error(),
			})
			return
		}
		if err := p.service.Delete(id); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	}
}
