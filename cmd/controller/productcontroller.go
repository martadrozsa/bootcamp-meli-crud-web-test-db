package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/product/domain"
	"net/http"
	"strconv"
)

type ProductController struct {
	service domain.ProductService
}

func CreateProductController(prodService domain.ProductService) *ProductController {
	return &(ProductController{service: prodService})
}

type requestProductPost struct {
	Name        string  `json:"name"  binding:"required"`
	ProductType string  `db:"product_type" json:"product_type" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type requestProductPatch struct {
	Price float64 `json:"price" binding:"required"`
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func (c *ProductController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		prod, err := c.service.GetById(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		ctx.JSON(http.StatusOK, prod)
	}
}

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productDto requestProductPost
		if err := ctx.ShouldBindJSON(&productDto); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"message": "Invalid input. Check the data entered",
				})
			return
		}
		newProduct, err := c.service.Create(ctx.Request.Context(), productDto.Name, productDto.ProductType, productDto.Description, productDto.Quantity, productDto.Price)

		if err != nil {
			ctx.JSON(http.StatusConflict, err)
			return
		}
		ctx.JSON(http.StatusCreated, newProduct)
	}
}

func (c *ProductController) UpdatePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		var productDto requestProductPatch

		if err := ctx.ShouldBindJSON(&productDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if productDto.Price <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
			return
		}

		err = c.service.UpdatePrice(ctx.Request.Context(), id, productDto.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		err = c.service.Delete(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, err)
	}
}
