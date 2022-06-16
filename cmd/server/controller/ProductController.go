package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/domain/product"
	"net/http"
	"strconv"
)

type ProductController struct {
	service product.ProductService
}

func CreateProductController(prodService product.ProductService) *ProductController {
	return &(ProductController{service: prodService})
}

type requestProductPost struct {
	Name        string  `json:"name"  binding:"required"`
	ProductType string  `db:"product_type" json:"product_type" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
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

		// transforma  o body  da requisição de json para o DTO
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
	return func(ctx *gin.Context) {}
	//TODO implement me
	panic("implement me")
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
	//TODO implement me
	panic("implement me")
}
