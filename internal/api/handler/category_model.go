package handler

import (
	"bakery-api/internal/dependency"
	"bakery-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryModelHandler struct {
	usecase *usecase.CategoryUseCase
}

func NewCategoryModelHandler() *CategoryModelHandler {
	return &CategoryModelHandler{
		usecase: dependency.GetCategoryUseCase(),
	}
}

func (h *CategoryModelHandler) CreateCategory(c *gin.Context) {
	Create(c, h.usecase.Create)
}

func (h *CategoryModelHandler) UpdateCategory(c *gin.Context) {
	Update(c, h.usecase.Update)
}

func (h *CategoryModelHandler) DeleteCategory(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

func (h *CategoryModelHandler) GetCategory(c *gin.Context) {
	FindById(c, h.usecase.FindById)
}
