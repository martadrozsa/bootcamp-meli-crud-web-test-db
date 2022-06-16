package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/cmd/server/controller"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/config"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/modules/repository"
	"github.com/martadrozsa/bootcamp-meli-crud-web-test/internal/modules/service"
)

func main() {

	db := config.ConnectDb()
	defer db.Close()

	router := gin.Default()
	groupV1 := router.Group("api/v1")

	productRepository := repository.CreateProductRepository(db)

	productService := service.CreateProductService(productRepository)
	productController := controller.CreateProductController(productService)

	productGroup := groupV1.Group("/products")
	productGroup.GET("/", productController.GetAll())
	productGroup.GET("/:id", productController.GetById())
	productGroup.POST("/", productController.Create())
	productGroup.PATCH("/:id", productController.UpdatePrice())
	productGroup.DELETE("/:id", productController.Delete())

	router.Run()

}
