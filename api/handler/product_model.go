package handler

import (
	"bakery-api/dependency"
	"bakery-api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductModelHandler struct {
	usecase *usecase.ProductUseCase
}

func NewProductModelHandler() *ProductModelHandler {
	return &ProductModelHandler{
		usecase: dependency.GetProductUseCase(),
	}
}

func (h *ProductModelHandler) CreateProduct(c *gin.Context) {
	Create(c, h.usecase.Create)
}

func (h *ProductModelHandler) UpdateProduct(c *gin.Context) {
	Update(c, h.usecase.Update)
}

func (h *ProductModelHandler) DeleteProduct(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

func (h *ProductModelHandler) GetProductById(c *gin.Context) {
	FindById(c, h.usecase.FindById)
}
