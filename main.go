package main

import (
	"fmt"
	"net/http"

	"epicnotification/epicnotification/routes"
)

func main() {
	// definitions
	port := ":8080"

	// logs
	fmt.Println("Running in http://localhost:8080")
	fmt.Println("available routes: http://localhost:8080/scrape/epicgames")

	// routes
	http.HandleFunc("/", routes.HomeRoute)
	http.HandleFunc("/scrape/epicgames", routes.EpicGamesRoute)
	// http.HandleFunc("/scrape/wiki", routes.Wiki)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("ERRROOROROR", err)
	}
}
