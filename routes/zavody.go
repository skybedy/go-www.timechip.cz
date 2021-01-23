package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-www.timechip.cz/conf"
	"go-www.timechip.cz/utils"
)

type ZavodType []map[string]string

func ZavodyJson(RaceYear string) ZavodType {
	var ZavodyMap ZavodType
	url := conf.API_TIMECHIP + "/homepage/zavody/" + RaceYear
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &ZavodyMap)
	return ZavodyMap
}

func Zavody(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.ExecuteTemplate(w, "zavody.html", struct {
		Title string
		Data  ZavodType
	}{
		Title: "Hlavní strana",
		Data:  ZavodyJson(vars["race-year"]),
	})
}
