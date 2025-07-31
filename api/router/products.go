package router

import (
	"bakery-api/api/handler"

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
	handler := handler.NewSizeModelHandler()

	r.POST("/", handler.CreateSize)
	r.PUT("/:id", handler.UpdateSize)
	r.DELETE("/:id", handler.DeleteSize)
	r.GET("/:id", handler.GetSizeById)
}

func Products(r *gin.RouterGroup) {
	h := handler.NewProductModelHandler()

	r.POST("/", h.CreateProduct)
	r.PUT("/:id", h.UpdateProduct)
	r.DELETE("/:id", h.DeleteProduct)
	r.GET("/:id", h.GetProductById)
}
