package error_handler

import (
	appconstant "bakery-api/app-constant"
	customerrors "bakery-api/configs/custom-errors"
	"bakery-api/internal/usecase/dto"
	"errors"
	"fmt"
	"net/http"

	customvalidator "bakery-api/internal/usecase/validator"

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
	var notFoundError customerrors.NotFoundError

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
	apiError := dto.NewAPIError(appconstant.INTERNAL_ERROR, "internal error", err.Error())
	ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse(http.StatusInternalServerError, apiError))
}

func handleValidationErrors(ctx *gin.Context, validationErrors *validator.ValidationErrors) {
	errorDetails := customvalidator.GetValidateError(validationErrors)
	apiError := dto.NewAPIError(appconstant.VALIDATION_ERROR, "binding error", errorDetails)
	ctx.JSON(http.StatusBadRequest, dto.ErrorResponse(http.StatusBadRequest, apiError))
}

func handleNotFoundError(ctx *gin.Context, notFoundError customerrors.NotFoundError) {
	defaultMessage := "Can not found %s"

	apiErrors := dto.NewAPIError(appconstant.NOT_FOUND_ERROR,
		fmt.Sprintf(defaultMessage, notFoundError.Message), nil)

	ctx.JSON(http.StatusNotFound, dto.ErrorResponse(http.StatusNotFound, apiErrors))
}
