package routes

import (
	"fmt"
	"net/http"
	"reflect"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"
	"epicnotification/epicnotification/utils/epicgames"
)

func EpicGamesRoute(w http.ResponseWriter, r *http.Request) {
	var url = utils.BackendURLs["epic-games"]
	res, err := scrape.EpicGamesScraper(url) // get response as string
	if err != nil {
		fmt.Println(
			fmt.Errorf("error in scraper response: %s", err),
		)
	}

	// do extract
	extracted := epicgames.ExtractData(res)
	fmt.Println("type of extracted", reflect.TypeOf(extracted))

	jsonData, err := utils.ToJSON(extracted)
	if err != nil {
		http.Error(w, "Error while converting response to json", http.StatusInternalServerError)
		return
	}
	fmt.Println("type of jsonData", reflect.TypeOf(jsonData))
	w.Write(jsonData)
}
