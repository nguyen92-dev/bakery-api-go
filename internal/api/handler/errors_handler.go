package handler

import (
	appconstant "bakery-api/app-constant"
	"bakery-api/internal/usecase/dto"
	validator2 "bakery-api/internal/usecase/validator"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	var validationErrors validator.ValidationErrors
	switch {
	case errors.As(err, &validationErrors):
		handlerValidationError(ctx, &validationErrors)
		return true
	default:
		ctx.JSON(500, dto.NewAPIResponse[any](nil, []dto.APIError{
			{
				Code:    appconstant.INTERNAL_ERROR,
				Message: "internal_error",
			},
		}))
		return true
	}
}

func handlerValidationError(ctx *gin.Context, err *validator.ValidationErrors) {
	apiError := validator2.GetValidateError(err)
	ctx.JSON(400, dto.NewAPIResponse[any](nil, apiError))
}
