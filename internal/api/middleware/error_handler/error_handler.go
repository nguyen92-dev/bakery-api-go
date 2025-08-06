package error_handler

import (
	custom_errors "bakery-api/configs/custom-errors"
	"bakery-api/internal/usecase/dto"
	"errors"
	"net/http"

	custom_validator "bakery-api/internal/usecase/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ThrowError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	c.Error(err)
	c.Abort()
	return true
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			handleError(ctx, err)
		}
	}
}

func handleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	var bindingErrors validator.ValidationErrors
	var notFoundError custom_errors.NotFoundError

	switch {
	case errors.As(err, &bindingErrors):
		handleValidationErrors(ctx, &bindingErrors)
	case errors.As(err, &notFoundError):
		handleNotFoundError(ctx, notFoundError)
	default:
		handleInternalError(ctx, err)
	}
}

func handleInternalError(ctx *gin.Context, err error) {
	apiError := []dto.APIError{
		{
			Code:    "500",
			Message: err.Error(),
		},
	}
	ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(apiError))
}

func handleValidationErrors(ctx *gin.Context, validationErrors *validator.ValidationErrors) {
	apiErrors := custom_validator.GetValidateError(validationErrors)

	ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(apiErrors))
}

func handleNotFoundError(ctx *gin.Context, notFoundError custom_errors.NotFoundError) {
	defaultMessage := "Can not found "

	apiErrors := []dto.APIError{
		{
			Code:    "404",
			Message: defaultMessage + notFoundError.Message,
		},
	}

	ctx.JSON(http.StatusNotFound, dto.NewErrorResponse(apiErrors))
}
