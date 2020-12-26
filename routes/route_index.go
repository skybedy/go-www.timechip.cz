package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go-www.timechip.cz/conf"
	"go-www.timechip.cz/utils"
)

type Zavod struct {
	IDZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
	KodZavodu   string `json:"kod_zavodu"`
	DatumZavodu string `json:"datum_zavodu"`
	Web         string `json:"web"`
	Icon        string `json:"icon"`
}

/*
type coinsData struct {
	IdZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
}*/

func NejblizsiZavody() map[string][]Zavod {

	var xyz map[string][]Zavod

	url := conf.API_TIMECHIP + "/homepage/nejblizsi-zavody"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//var zavody []Zavod
	err = json.Unmarshal(body, &xyz)
	fmt.Println(xyz["nejblizsi_zavody"])

	return xyz
}

func index(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Data  map[string][]Zavod
	}{
		Title: "Hlavní strana",
		Data:  NejblizsiZavody(),
	})
}
