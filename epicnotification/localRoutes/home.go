package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func HomeRoute(w http.ResponseWriter, r *http.Request) {
func HomeRoute(c *gin.Context) {
	scrapeRoutes := map[string]interface{}{
		"EpicGames": "http://localhost:8080/scrape/epicgames",
		"Wikipedia": "http://localhost:8080/scrape/wiki",
	}
	routes := map[string]interface{}{
		"home":   "http://localhost:8080",
		"scrape": scrapeRoutes,
	}

	response, err := json.Marshal(routes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Json'a donusturme patladi"})
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", response)
}
