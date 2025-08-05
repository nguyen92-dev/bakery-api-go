package web_api

import (
	"bakery-api/configs"
	"bakery-api/internal/api"
	"bakery-api/internal/infra/persisstence/database"
)

type WebAPI struct {
}

func NewWebAPI() *WebAPI {
	return &WebAPI{}
}

func (app *WebAPI) Run() {
	configs.Cfg = configs.GetConfig()

	err := database.InitDb(configs.Cfg)
	if err != nil {
		panic(err)
	}
	api.InitServer()
}
