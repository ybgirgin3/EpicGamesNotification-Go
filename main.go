package main

import (
	"fmt"

	"epicnotification/epicnotification/apps/scrape"
)

func main() {

	var url = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=TR&allowCountries=TR"
	// fmt.Println("headers:", reflect.TypeOf(headers))
	scrapeResult, scraperError := scrape.Scraper(url)
	if scraperError != nil {
		err := fmt.Errorf("error in scraper response: %s", scraperError)
		fmt.Println(err)
	}
	fmt.Println(*scrapeResult)
}
