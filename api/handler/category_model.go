package handler

import (
	"bakery-api/usecase"
	"bakery-api/usecase/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryModelHanldler struct {
	usecase *usecase.CategoryUseCase
}

func NewCategoryModelHandler(usecase *usecase.CategoryUseCase) *CategoryModelHanldler {
	return &CategoryModelHanldler{
		usecase: usecase,
	}
}

func (h *CategoryModelHanldler) CreateCategory(c *gin.Context) {
	request := new(dto.CategoryRequestDto)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := h.usecase.Create(c, *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, response)
}

func (h *CategoryModelHanldler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	request := new(dto.CategoryRequestDto)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := h.usecase.Update(c, id, *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}

func (h *CategoryModelHanldler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.usecase.Delete(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204) // No Content
}

func (h *CategoryModelHanldler) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	response, err := h.usecase.FindById(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
