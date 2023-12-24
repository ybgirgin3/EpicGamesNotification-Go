package routes

import (
	"encoding/csv"
	"net/http"
	"os"

	"github.com/gocolly/colly"
)

type Scrape struct {
	url string
}

type PokemonProduct struct {
	url, image, name, price string
}

func (s Scrape) Wiki(w http.ResponseWriter, r *http.Request) {
	// definitions
	var url = s.url
	c := colly.NewCollector()
	var pokemonProducts []PokemonProduct

	// visit site
	c.Visit(url)

	/*
			OnRequest(): Called before performing an HTTP request with Visit().
			OnError(): Called if an error occurred during the HTTP request.
			OnResponse(): Called after receiving a response from the server.
			OnHTML(): Called right after OnResponse() if the received content is HTML.
			OnScraped(): Called after all OnHTML() callback executions.


		// default callbacks
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL)
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong: ", err)
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Page visited: ", r.Request.URL)
		})

		c.OnHTML("a", func(e *colly.HTMLElement) {
			// printing all URLs associated with the a links in the page
			fmt.Println("%v", e.Attr("href"))
		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println(r.Request.URL, " scraped!")
		})
	*/
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initialize a new product instance
		pokemonProduct := PokemonProduct{}

		// scraping the data of interest
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.name = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	file, err := os.Create("pokemons.csv")
	if err != nil {
		http.Error(w, "Error while converting response to json", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// initalizing a file writer
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{
		"url", "image", "name", "price",
	}

	// writing the column headers
	writer.Write(headers)

	// adding each product to csv output file
	for _, pr := range pokemonProducts {
		record := []string{
			pr.url, pr.image, pr.name, pr.price,
		}

		// write a new CSV record
		writer.Write(record)
	}

	defer writer.Flush()

}
