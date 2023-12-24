package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	urls "epicnotification/epicnotification/utils"
	"epicnotification/epicnotification/utils/epicgames"
)

func EpicGamesRoute(w http.ResponseWriter, r *http.Request) {
	var url = urls.BackendURLs["epic-games"]
	res, err := scrape.EpicGamesScraper(url) // get response as string
	if err != nil {
		fmt.Println(
			fmt.Errorf("error in scraper response: %s", err),
		)
	}

	// do extract
	extracted := epicgames.ExtractData(res)

	jsonData, err := json.Marshal(extracted)
	if err != nil {
		http.Error(w, "Error while converting response to json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
