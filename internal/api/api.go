package api

import (
	"bakery-api/internal/api/middleware/error_handler"
	"bakery-api/internal/api/router"
	customValidator "bakery-api/internal/usecase/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	r := gin.Default()

	r.Use(error_handler.ErrorHandler())

	RegisterRoutes(r)
	RegisterValidator()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		categories := v1.Group("/categories")
		router.Categories(categories)

		sizes := v1.Group("/sizes")
		router.Sizes(sizes)

		product := v1.Group("/products")
		router.Products(product)
	}
}

func RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("non_special_char", customValidator.ValidateNonSpecialCharacter)
		if err != nil {
			panic(err)
		}
	}
}
