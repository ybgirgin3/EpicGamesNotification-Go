package routes

import (
	"fmt"
	"net/http"
	"reflect"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"
	"epicnotification/epicnotification/utils/epicgames"

	"github.com/gin-gonic/gin"
)

// func EpicGamesRoute(w http.ResponseWriter, r *http.Request) {
func EpicGamesRoute(c *gin.Context) {
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

	response, err := utils.ToJSON(extracted)
	if err != nil {
		// http.Error(w, "Error while converting response to json", http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while converting response to json in epicgames route"})
		return
	}
	fmt.Println("type of jsonData", reflect.TypeOf(response))
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
