package handler

import (
	"bakery-api/configs"
	"bakery-api/dependency"
	"bakery-api/usecase"

	"github.com/gin-gonic/gin"
)

type SizeModelHandler struct {
	usecase *usecase.SizeUseCase
}

func NewSizeModelHandler(cfg *configs.Config) *SizeModelHandler {
	return &SizeModelHandler{
		usecase: usecase.NewSizeUseCase(cfg, dependency.GetSizeRepository(cfg)),
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
