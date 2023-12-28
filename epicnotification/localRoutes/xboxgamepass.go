package routes

import (
	"net/http"

	"epicnotification/epicnotification/apps/scrape"
	"epicnotification/epicnotification/utils"

	"github.com/gin-gonic/gin"
)

func XBoxGamePassRoute(c *gin.Context) {
	url := "https://www.xbox.com/en-US/xbox-game-pass/games"
	res, err := scrape.XboxGamePassScraper(url)

	if err != nil {
		panic(err)
	}

	// convert to json
	response, err := utils.ToJSON(res)
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
