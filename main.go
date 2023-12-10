package main

import (
	"fmt"

	"epicnotification/epicnotification/apps/scrape"
)

func main() {
	var url = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=TR&allowCountries=TR"
	// var headers = scrape.RequestHeaders
	headers := `{
        "accept": "application/json, text/plain, */*",
        "accept-language": "en-US,en;q=0.9",
        "sec-ch-ua": '"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "empty",
        "sec-fetch-mode": "cors",
        "sec-fetch-site": "same-site",
        "x-requested-with": "XMLHttpRequest",
        "Referer": "https://store.epicgames.com/en-US/free-games",
        "Referrer-Policy": "no-referrer-when-downgrade",
    }`
	// fmt.Println("headers:", reflect.TypeOf(headers))
	scrapeResult, scraperError := scrape.Scraper(url, headers)
	if scraperError != nil {
		err := fmt.Errorf("Error in scraper response: %s", scraperError)
		fmt.Println(err)
	}
	fmt.Println(scrapeResult)

}
