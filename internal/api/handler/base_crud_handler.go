package handler

import (
	appconstant "bakery-api/app-constant"
	"bakery-api/internal/usecase/dto"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create[TRequest any, TResponse any](c *gin.Context,
	usecaseCreate func(ctx context.Context, request TRequest) (TResponse, error)) {
	request := new(TRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, BadRequestError(err, appconstant.INVALID_INPUT))
		return
	}

	response, err := usecaseCreate(c, *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	apiResponse := dto.NewAPIResponse(response, nil)
	c.JSON(201, apiResponse)
}

func Update[TRequest any, TResponse any](c *gin.Context,
	usecaseUpdate func(ctx context.Context, id uint, request TRequest) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	request := new(TRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := usecaseUpdate(c, uint(id), *request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}

func Delete(c *gin.Context,
	usecaseDelete func(ctx context.Context, id uint) error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := usecaseDelete(c, uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "Deleted successfully"})
}

func FindById[TResponse any](c *gin.Context,
	usecaseFindById func(ctx context.Context, id uint) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	response, err := usecaseFindById(c, uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}

func BadRequestError(err error, code string) dto.APIResponse[any] {
	apiError := dto.NewAPIError(code, err.Error())
	return dto.NewAPIResponse[any](nil, apiError)
}
