package router

import (
	"bakery-api/api/handler"

	"github.com/gin-gonic/gin"
)

func Categories(r *gin.RouterGroup, handler *handler.CategoryModelHanldler) {
	h := handler

	r.POST("/", h.CreateCategory)
	r.PUT("/:id", h.UpdateCategory)
	r.DELETE("/:id", h.DeleteCategory)
	r.GET("/:id", h.GetCategory)
}
