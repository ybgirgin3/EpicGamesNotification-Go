package scrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type Scrape struct {
	url     string
	headers string
}

func Scraper(url string, headers string) (string, error) {
	scraperObject := Scrape{url, headers}
	// fmt.Println("s in Scraper fonk", scraperObject)
	result, err := scraperObject.PerformScrape()
	return result, err

}

func (s Scrape) PerformScrape() (string, error) {
	// create request object
	// var headers string
	var url = s.url

	// create request object
	headers, err := createHeaders()
	if err != nil {
		err := fmt.Errorf("unable to create headers object: %s", err)
		return "", err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := fmt.Errorf("unable to create request object: %s", err)
		return "", err
	}

	request.Header = headers

	// do request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err := fmt.Errorf("unable to send request: %s", err)
		return "", err
	}
	defer response.Body.Close()

	// control status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("unexpected status code: %v", response.Status)
	}

	// read body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body :%v", err)
	}
	result := string(body)
	return result, nil

}

func createHeaders() (http.Header, error) {
	var headerMap map[string]string

	fmt.Println("requestheaders in createheaders:", RequestHeaders, reflect.TypeOf(RequestHeaders))

	err := json.Unmarshal([]byte(RequestHeaders), &headerMap)
	if err == nil {
		// return nil, fmt.Errorf("error parsing json to headerset", err)
		fmt.Println("")
	}

	fmt.Println("headermap in createheaders", headerMap)

	// make headers
	headers := make(http.Header)
	for key, value := range headerMap {
		headers.Add(key, value)
	}
	// fmt.Println("headers in createheaders", headers)
	return headers, nil
}
