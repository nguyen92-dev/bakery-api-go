package handler

import (
	"bakery-api/internal/api/middleware/error_handler"
	"bakery-api/internal/usecase/dto"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create[TRequest any, TResponse any](c *gin.Context,
	useCaseCreate func(ctx context.Context, request TRequest) (TResponse, error)) {
	request := new(TRequest)
	if error_handler.ThrowError(c, c.ShouldBindJSON(request)) {
		return
	}

	response, err := useCaseCreate(c, *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	apiResponse := dto.SuccessResponse(http.StatusCreated, response)
	c.JSON(201, apiResponse)
}

func Update[TRequest any, TResponse any](c *gin.Context,
	useCaseUpdate func(ctx context.Context, id uint, request TRequest) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	request := new(TRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		ErrorHandler(c, err)
		return
	}

	response, err := useCaseUpdate(c, uint(id), *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dto.SuccessResponse(http.StatusOK, response))
}

func Delete(c *gin.Context,
	useCaseDelete func(ctx context.Context, id uint) error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := useCaseDelete(c, uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "Deleted successfully"})
}

func FindById[TResponse any](c *gin.Context,
	useCaseFindById func(ctx context.Context, id uint) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	response, err := useCaseFindById(c, uint(id))
	if error_handler.ThrowError(c, err) {
		return
	}
	c.JSON(200, dto.SuccessResponse(http.StatusOK, response))
}
