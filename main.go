package main

import (
	"go-www.timechip.cz/routes"
	"go-www.timechip.cz/utils"
)

func main() {
	router := routes.NewRouter()
	utils.LoadTemplates("/var/www/timechip.cz/go-www.timechip.cz/templates/*.html")
	utils.HttpServer(router)
}
