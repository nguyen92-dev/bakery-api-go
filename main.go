package main

import webapi "bakery-api/cmd/web-api"

func main() {
	app := webapi.NewWebAPI()
	app.Run()
}
