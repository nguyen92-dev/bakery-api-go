package handler

import (
	appconstant "bakery-api/app-constant"
	"bakery-api/internal/usecase/dto"
	validator2 "bakery-api/internal/usecase/validator"
	"errors"
	"net/http"

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
		internalError := dto.NewAPIError(appconstant.InternalError, "internal error", err.Error())
		ctx.JSON(500, dto.ErrorResponse(http.StatusInternalServerError, internalError))
		return true
	}
}

func handlerValidationError(ctx *gin.Context, err *validator.ValidationErrors) {
	details := validator2.GetValidateError(err)
	varlidationError := dto.NewAPIError(appconstant.ValidationError, "validation error", details)
	ctx.JSON(400, dto.ErrorResponse(http.StatusBadRequest, varlidationError))
}
