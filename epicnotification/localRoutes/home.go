package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Context) {
	scrapeRoutesFunc := func(baseUrl string) interface{} {
		return map[string]interface{}{
			"EpicGames": baseUrl + "/scrape/epic-games",
			"Wikipedia": baseUrl + "/scrape/wiki",
			"Xbox":      baseUrl + "/scrape/xbox-gamepass",
		}
	}

	base := "http://localhost:8080"

	routes := map[string]interface{}{
		"home":   base,
		"scrape": scrapeRoutesFunc(base),
	}

	response, err := json.Marshal(routes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Json'a donusturme patladi"})
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
