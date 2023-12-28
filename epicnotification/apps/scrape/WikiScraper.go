package scrape

import (
	"github.com/gocolly/colly"
)

type WS struct {
	url string
}

func WikiScraper(url string) (map[string]string, error) {
	so := WS{url}
	result, err := so.PerformScrape()
	return result, err
}

func (s WS) PerformScrape() (map[string]string, error) {

	c := colly.NewCollector()

	links := make(map[string]string)

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.ChildAttr("a", "href")

		links[title] = link
	})

	err := c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")
	if err != nil {
		panic(err)
	}

	return links, nil

}
