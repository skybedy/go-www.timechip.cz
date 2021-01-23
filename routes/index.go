package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"go-www.timechip.cz/conf"
	"go-www.timechip.cz/utils"
)

type NejblizsiZavod struct {
	IDZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
	KodZavodu   string `json:"kod_zavodu"`
	DatumZavodu string `json:"datum_zavodu"`
	Web         string `json:"web"`
	Icon        string `json:"icon"`
}

func NejblizsiZavody(CurrentYear string) map[string][]NejblizsiZavod {
	var xyz map[string][]NejblizsiZavod
	url := conf.API_TIMECHIP + "/homepage/nejblizsi-zavody/" + CurrentYear
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &xyz)
	return xyz
}

func Index(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Data  map[string][]NejblizsiZavod
	}{
		Title: "Hlavní strana",
		Data:  NejblizsiZavody(strconv.Itoa(t.Year())),
	})
}
