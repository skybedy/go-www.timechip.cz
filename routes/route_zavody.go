package routes

import (
	"net/http"

	"go-www.timechip.cz/utils"
)

func Zavody(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Data  map[string][]Zavod
	}{
		Title: "Hlavní strana",
		Data:  NejblizsiZavody(),
	})
}
