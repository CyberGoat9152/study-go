package main

import (
	"fmt"
	"restapi/1/cmd/controller"
	"restapi/1/cmd/db"
	"restapi/1/cmd/repository"
	"restapi/1/cmd/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//db connection
	fmt.Println("iniciando")
	dbConn, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConn)
	fmt.Println("connectou")
	//model controller
	ProductUsecase := usecase.NewProductUserCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
