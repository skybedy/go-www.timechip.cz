package routes

import "net/http"

func Vysledky(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://vysledky.timechip.cz", 301)
}
