package handler

import (
	"bakery-api/dependency"
	"bakery-api/usecase"

	"github.com/gin-gonic/gin"
)

type SizeModelHandler struct {
	usecase *usecase.SizeUseCase
}

func NewSizeModelHandler() *SizeModelHandler {
	return &SizeModelHandler{
		usecase: usecase.NewSizeUseCase(dependency.GetSizeRepository(), dependency.GetCategoryRepository()),
	}
}

func (h *SizeModelHandler) CreateSize(c *gin.Context) {
	Create(c, h.usecase.Create)
}

func (h *SizeModelHandler) UpdateSize(c *gin.Context) {
	Update(c, h.usecase.Update)
}

func (h *SizeModelHandler) DeleteSize(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

func (h *SizeModelHandler) GetSizeById(c *gin.Context) {
	FindById(c, h.usecase.FindById)
}
