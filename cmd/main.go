package main

import (
	"bakery-api/api"
	"bakery-api/configs"
	"bakery-api/infra/persisstence/database"
)

func main() {
	configs.Cfg = configs.GetConfig()

	err := database.InitDb(configs.Cfg)
	if err != nil {
		panic(err)
	}
	api.InitServer()
}
