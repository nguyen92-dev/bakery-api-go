package api

import (
	"bakery-api/api/handler"
	"bakery-api/api/router"
	"bakery-api/configs"
	"bakery-api/domain/model"
	"bakery-api/infra/persisstence/database"
	"bakery-api/infra/persisstence/repository"
	"bakery-api/usecase"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *configs.Config) {
	r := gin.Default()
	preloads := []database.PreloadEntity{}
	repository := repository.NewBaseRepository[model.Category](cfg, preloads)
	usecase := usecase.NewCategoryUseCase(repository)
	handler := handler.NewCategoryModelHandler(usecase)
	RegisterRoutes(r, handler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func RegisterRoutes(r *gin.Engine, hanlder *handler.CategoryModelHanldler) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		categories := v1.Group("/categories")

		router.Categories(categories, hanlder)
	}
}
