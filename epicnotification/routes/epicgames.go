package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"
)

func EpicGamesRoute(w http.ResponseWriter, r *http.Request) {
	var url = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=TR&allowCountries=TR"
	res, err := scrape.EpicGamesScraper(url) // get response as string
	if err != nil {
		fmt.Println(
			fmt.Errorf("error in scraper response: %s", err),
		)
	}

	// do extract
	extracted := utils.ExtractData(res)

	jsonData, err := json.Marshal(extracted)
	if err != nil {
		http.Error(w, "Error while converting response to json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
