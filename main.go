package main

import (
	"fmt"

	localRoutes "epicnotification/epicnotification/localRoutes"

	"github.com/gin-gonic/gin"
)

func main() {
	// definitions

	// logs
	fmt.Println("Running in http://localhost:8080")
	fmt.Println("available routes: http://localhost:8080/scrape/epicgames")
	fmt.Println("available routes: http://localhost:8080/scrape/wiki")

	// router
	routes := gin.Default()

	// routes
	routes.GET("/", localRoutes.HomeRoute)
	routes.GET("/scrape/:retailer", handlerSelector)

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

func handlerSelector(c *gin.Context) {
	retailer := c.Param("retailer")
	if retailer == "epicgames" {
		localRoutes.EpicGamesRoute(c)
	} else if retailer == "wiki" {
		localRoutes.WikiRoute(c)
	}

}
