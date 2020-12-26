package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"go-www.timechip.cz/routes"
	"go-www.timechip.cz/utils"
)

const Port = "1301"

var templates *template.Template

type Zavod struct {
	IDZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
	KodZavodu   string `json:"kod_zavodu"`
	DatumZavodu string `json:"datum_zavodu"`
	Icon        string `json:"icon"`
}

/*
type coinsData struct {
	IdZavodu    string `json:"id_zavodu"`
	NazevZavodu string `json:"nazev_zavodu"`
}*/

func NejblizsiZavody() {
	url := "http://localhost:1313/homepage/nejblizsi-zavody"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	if err != nil {
		log.Fatal(err)
	}

	//textBytes := []byte(body)

	//zavody := Zavod{}
	var zavody []Zavod

	//err := json.Unmarshal(textBytes, &zavody)
	err = json.Unmarshal(body, &zavody)

	/*
		var c []coinsData
		err = json.Unmarshal(body, &c)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			fmt.Println(err)
			return
		}*/

	//fmt.Println(zavody.IDzavodu)

	for k := range zavody {
		fmt.Println(zavody[k].NazevZavodu)

	}
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func index(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:1313/homepage/nejblizsi-zavody"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var zavody []Zavod

	err = json.Unmarshal(body, &zavody)

	for k := range zavody {
		fmt.Println(zavody[k].NazevZavodu)
	}

	ExecuteTemplate(w, "index.html", struct {
		Title  string
		Zavody []Zavod
	}{
		Title:  "Hlavní strana",
		Zavody: zavody,
	})

}

func main() {
	router := routes.NewRouter()
	utils.LoadTemplates("templates/*.html")
	utils.HttpServer(router)
}
