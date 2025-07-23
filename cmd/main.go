package main

import (
	"bakery-api/api"
	"bakery-api/configs"
	"bakery-api/infra/persisstence/database"
)

func main() {
	cfg := configs.GetConfig()

	err := database.InitDb(cfg)
	if err != nil {
		panic(err)
	}
	api.InitServer(cfg)
}
