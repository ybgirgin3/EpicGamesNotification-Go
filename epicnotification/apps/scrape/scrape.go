package scrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Scrape struct {
	url     string
	headers string
}

func Scraper(url string) (*ResponseStruct, error) {
	scraperObject := Scrape{url, RequestHeaders}
	result, err := scraperObject.PerformScrape()
	return result, err

}

func (s Scrape) PerformScrape() (*ResponseStruct, error) {
	// create request object
	// var headers string
	var url = s.url
	var responseData ResponseStruct

	// create request object
	headers, err := createHeaders()
	if err != nil {
		err := fmt.Errorf("unable to create headers object: %s", err)
		return nil, err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := fmt.Errorf("unable to create request object: %s", err)
		return nil, err
	}

	request.Header = headers

	// do request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err := fmt.Errorf("unable to send request: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	// control status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %v", response.Status)
	}

	// read body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body :%v", err)
	}
	_result := string(body)
	fmt.Println("_result", _result)

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, fmt.Errorf("error exporting response body :%v", err)
	}

	return &responseData, nil

}

func createHeaders() (http.Header, error) {
	var headerMap map[string]string

	err := json.Unmarshal([]byte(RequestHeaders), &headerMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing json to headerset", err)
	}

	// make headers
	headers := make(http.Header)
	for key, value := range headerMap {
		headers.Add(key, value)
	}
	return headers, nil
}
