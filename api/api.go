package api

import (
	"bakery-api/api/router"
	"bakery-api/configs"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *configs.Config) {
	r := gin.Default()
	RegisterRoutes(r, cfg)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *configs.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		categories := v1.Group("/categories")

		router.Categories(categories, cfg)
	}
}
