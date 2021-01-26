package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-www.timechip.cz/routes"
	"go-www.timechip.cz/utils"
)

func main() {
	//router := routes.NewRouter()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.Index).Methods("GET")
	router.HandleFunc("/zavody/{race-year}", routes.Zavody).Methods("GET")
	router.HandleFunc("/vysledky/{race-year}", routes.Vysledky).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	utils.LoadTemplates("templates/*.html")
	utils.HttpServer(router)
}
