package handler

import (
	customerrors "bakery-api/configs/custom-errors"
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
	if error_handler.ThrowError(c, err) {
		return
	}
	apiResponse := dto.SuccessResponse(http.StatusCreated, response)
	c.JSON(201, apiResponse)
}

func Update[TRequest any, TResponse any](c *gin.Context,
	useCaseUpdate func(ctx context.Context, id uint, request TRequest) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		error_handler.ThrowError(c, customerrors.InvalidIdError())
		return
	}
	request := new(TRequest)
	if error_handler.ThrowError(c, c.ShouldBindJSON(request)) {
		return
	}

	response, err := useCaseUpdate(c, uint(id), *request)
	if error_handler.ThrowError(c, err) {
		return
	}
	c.JSON(200, dto.SuccessResponse(http.StatusOK, response))
}

func Delete(c *gin.Context,
	useCaseDelete func(ctx context.Context, id uint) error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		error_handler.ThrowError(c, customerrors.InvalidIdError())
		return
	}
	if error_handler.ThrowError(c, useCaseDelete(c, uint(id))) {
		return
	}
	c.JSON(204, gin.H{"message": "Deleted successfully"})
}

func FindById[TResponse any](c *gin.Context,
	useCaseFindById func(ctx context.Context, id uint) (TResponse, error)) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id <= 0 {
		error_handler.ThrowError(c, customerrors.InvalidIdError())
		return
	}
	response, err := useCaseFindById(c, uint(id))
	if error_handler.ThrowError(c, err) {
		return
	}
	c.JSON(200, dto.SuccessResponse(http.StatusOK, response))
}

//TODO: Implement pagelist with dynamic filter
