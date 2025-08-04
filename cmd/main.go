package main

import (
	"bakery-api/configs"
	"bakery-api/internal/api"
	"bakery-api/internal/infra/persisstence/database"
)

func main() {
	configs.Cfg = configs.GetConfig()

	err := database.InitDb(configs.Cfg)
	if err != nil {
		panic(err)
	}
	api.InitServer()
}
