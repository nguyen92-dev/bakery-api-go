package router

import (
	"bakery-api/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func Categories(r *gin.RouterGroup) {
	h := handler.NewCategoryModelHandler()

	r.POST("/", h.CreateCategory)
	r.PUT("/:id", h.UpdateCategory)
	r.DELETE("/:id", h.DeleteCategory)
	r.GET("/:id", h.GetCategory)
}

func Sizes(r *gin.RouterGroup) {
	h := handler.NewSizeModelHandler()

	r.POST("/", h.CreateSize)
	r.PUT("/:id", h.UpdateSize)
	r.DELETE("/:id", h.DeleteSize)
	r.GET("/:id", h.GetSizeById)
}

func Products(r *gin.RouterGroup) {
	h := handler.NewProductModelHandler()

	r.POST("/", h.CreateProduct)
	r.PUT("/:id", h.UpdateProduct)
	r.DELETE("/:id", h.DeleteProduct)
	r.GET("/:id", h.GetProductById)
}
