package handler

import (
	"bakery-api/configs"
	"bakery-api/dependency"
	"bakery-api/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryModelHanldler struct {
	usecase *usecase.CategoryUseCase
}

func NewCategoryModelHandler(cfg *configs.Config) *CategoryModelHanldler {
	return &CategoryModelHanldler{
		usecase: usecase.NewCategoryUseCase(cfg, dependency.GetCategoryRepository(cfg)),
	}
}

func (h *CategoryModelHanldler) CreateCategory(c *gin.Context) {
	Create(c, h.usecase.Create)
}

func (h *CategoryModelHanldler) UpdateCategory(c *gin.Context) {
	Update(c, h.usecase.Update)
}

func (h *CategoryModelHanldler) DeleteCategory(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

func (h *CategoryModelHanldler) GetCategory(c *gin.Context) {
	FindById(c, h.usecase.FindById)
}
