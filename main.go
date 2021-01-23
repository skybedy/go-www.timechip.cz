package main

import (
	"go-www.timechip.cz/routes"
	"go-www.timechip.cz/utils"
)

func main() {
	router := routes.NewRouter()
	utils.LoadTemplates("templates/*.html")
	utils.HttpServer(router)
}
