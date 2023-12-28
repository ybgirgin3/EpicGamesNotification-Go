package routes

import (
	"fmt"
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"

	"github.com/gin-gonic/gin"
)

func WikiRoute(c *gin.Context) {
	url := "https://scrapeme.live/shop/"
	res, err := scrape.WikiScraper(url)
	if err != nil {
		fmt.Println("unable to scrape wiki")
	}

	// convert to json
	response, err := utils.ToJSON(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to convert response to json in WikiRoute"})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
