package router

import (
	"bakery-api/api/handler"
	"bakery-api/configs"

	"github.com/gin-gonic/gin"
)

func Categories(r *gin.RouterGroup, cfg *configs.Config) {
	h := handler.NewCategoryModelHandler(cfg)

	r.POST("/", h.CreateCategory)
	r.PUT("/:id", h.UpdateCategory)
	r.DELETE("/:id", h.DeleteCategory)
	r.GET("/:id", h.GetCategory)
}

func Sizes(r *gin.RouterGroup, cfg *configs.Config) {
	handler := handler.NewSizeModelHandler(cfg)

	r.POST("/", handler.CreateSize)
	r.PUT("/:id", handler.UpdateSize)
	r.DELETE("/:id", handler.DeleteSize)
	r.GET("/:id", handler.GetSizeById)
}
