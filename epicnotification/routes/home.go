package routes

import (
	"encoding/json"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	scrapeRoutes := map[string]interface{}{
		"EpicGames": "http://localhost:8080/scrape/epicgames",
		"Wikipedia": "http://localhost:8080/scrape/wiki",
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
