package scrape

import (
	"github.com/gocolly/colly"
)

type XB struct {
	url string
}

func XboxGamePassScraper(url string) ([]string, error) {
	so := XB{url}
	response, err := so.PerformScrape()
	if err != nil {
		panic(err)

	}
	return response, nil
}

func (x XB) PerformScrape() ([]string, error) {
	c := colly.NewCollector()

	// links := make(map[string]string)
	links := []string{}

	c.OnHTML("section.m-product-placement-item", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a.gameDivLink", "href")
		links = append(links, link)
	})

	err := c.Visit(x.url)
	if err != nil {
		panic(err)
	}

	return links, nil

}
