package controller

import (
	"net/http"
	"restapi/1/cmd/model"
	"restapi/1/cmd/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id must be int, not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id must be int",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		response := model.Response{
			Message: "ops, something get wrong try again later",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	productId, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, productId)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id must be int, not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id must be int",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	var product model.Product
	errBody := ctx.BindJSON(&product)
	if errBody != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	errUpadte := p.productUsecase.UpdateProduct(productId, product)

	if errUpadte != nil {
		response := model.Response{
			Message: "ops, something get wrong try again later",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Successfully updated",
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id must be int, not null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id must be int",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	errDelete := p.productUsecase.DeleteProduct(productId)

	if errDelete != nil {
		response := model.Response{
			Message: "ops, something get wrong try again later",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.Response{
		Message: "Successfully deleted",
	}
	ctx.JSON(http.StatusOK, response)
}
