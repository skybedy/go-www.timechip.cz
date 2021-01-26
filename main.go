package main

import (
	"go-www.timechip.cz/routes"
	"go-www.timechip.cz/utils"
)

func main() {
	router := routes.NewRouter()

	//router := mux.NewRouter()
	//router.HandleFunc("/", routes.Index).Methods("GET")
	//router.HandleFunc("/api/vypis-roku/{pohlavi}", routes.VypisRoku).Methods("GET")
	//router.HandleFunc("/export", routes.Export).Methods("GET")
	//router.HandleFunc("/api/insert-to-db", routes.InsertToDB).Methods("POST")

	//path := "/var/dev/go/src/vianocezkrabicky.timechip.cz/"

	//staticFileDirectory := http.Dir("./static/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	//staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	//router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	utils.LoadTemplates("/var/www/timechip.cz/go-www.timechip.cz/templates/*.html")
	utils.HttpServer(router)

}
