package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"
)

// type T struct {
// 	HomeRoute string `json:"homeRoute"`
// 	SubRoute  string `json:"subRoute"`
// }

func main() {
	// definitions
	port := ":8080"

	// logs
	fmt.Println("Running in http://localhost:8080")
	fmt.Println("available routes: http://localhost:8080/scrape/epicgames")

	// routes
	http.HandleFunc("/", home)
	http.HandleFunc("/scrape/epicgames", scrapeEndpoint)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("ERRROOROROR", err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	scrapeRoutes := map[string]interface{}{
		"EpicGames": "http://localhost:8080/scrape/epicgames",
	}
	routes := map[string]interface{}{
		"home":   "http://localhost:8080",
		"scrape": scrapeRoutes,
	}

	jsonData, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, "Json'a donusturme patladi", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func scrapeEndpoint(w http.ResponseWriter, r *http.Request) {

	var url = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=TR&allowCountries=TR"
	res, err := scrape.Scraper(url) // get response as string
	if err != nil {
		fmt.Println(
			fmt.Errorf("error in scraper response: %s", err),
		)
	}

	// do extract
	extracted := utils.ExtractData(res)

	jsonData, err := json.Marshal(extracted)
	if err != nil {
		http.Error(w, "Error while converting response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}
