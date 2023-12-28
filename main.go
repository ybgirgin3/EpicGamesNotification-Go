package main

import (
	"fmt"
	"net/http"

	localRoutes "epicnotification/epicnotification/localRoutes"

	"github.com/gin-gonic/gin"
)

func main() {
	// definitions
	methods := make(map[string]func(c *gin.Context))

	methods["epic-games"] = localRoutes.EpicGamesRoute
	methods["wiki"] = localRoutes.WikiRoute
	methods["xbox-gamepass"] = localRoutes.XBoxGamePassRoute

	// logs
	baseURL := "http://localhost:8080"
	for index := range methods {
		fmt.Println("available routes: " + baseURL + "/scrape/" + index)
	}

	// router
	routes := gin.Default()

	// routes
	routes.GET("/", localRoutes.HomeRoute)
	routes.GET("/scrape/:retailer", func(c *gin.Context) {
		retailer := c.Param("retailer")

		// look for if retailer exists in maps
		// if exists pass context to handler
		if handler, ok := methods[retailer]; ok {
			handler(c)
		} else {
			c.Data(http.StatusNotFound, "Retailer Not Found", nil)
		}
	})

	port := ":8080"
	err := routes.Run(port)
	if err != nil {
		fmt.Println("ERRROOROROR", err)
		return
	}

	// routes
	// http.HandleFunc("/", routes.HomeRoute)
	// http.HandleFunc("/scrape/epicgames", routes.EpicGamesRoute)
	// http.HandleFunc("/scrape/wiki", routes.WikiRoute)
	//
	// err := http.ListenAndServe(port, nil)
	// if err != nil {
	// 	fmt.Println("ERRROOROROR", err)
	// }
}

// func handlerSelector(c *gin.Context) {
// 	retailer := c.Param("retailer")
// 	// if retailer == "epicgames" {
// 	// 	localRoutes.EpicGamesRoute(c)
// 	// } else if retailer == "wiki" {
// 	// 	localRoutes.WikiRoute(c)
// 	// }
//
// }
