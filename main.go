package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"
)

func main() {
	// definitions
	port := ":8080"

	// logs
	fmt.Println("Running in http://localhost:8080")
	fmt.Println("available routes: http://localhost:8080/scrape/epicgames")

	// routes
	http.HandleFunc("/scrape/epicgames", scrapeEndpoint)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("ERRROOROROR", err)
	}

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
