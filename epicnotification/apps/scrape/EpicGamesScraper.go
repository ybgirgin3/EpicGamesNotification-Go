package scrape

import (
	"fmt"
	"io"
	"net/http"
	"reflect"

	"epicnotification/epicnotification/helpers"
	"epicnotification/epicnotification/utils"
)

type Scrape struct {
	url     string
	headers string
}

func EpicGamesScraper(url string) (helpers.ApiResponse, error) {
	scraperObject := Scrape{url, helpers.EpicGamesRequestHeaders}
	result, err := scraperObject.PerformScrape()
	return result, err
}

func (s Scrape) PerformScrape() (helpers.ApiResponse, error) {
	// create request object
	var url = s.url
	var jsonHeaders = s.headers
	fmt.Println("TYPE OF HEADERS", reflect.TypeOf(jsonHeaders))

	// create request object
	headers, err := utils.CreateHeaders(jsonHeaders)
	if err != nil {
		err := fmt.Errorf("unable to create headers object: %s", err)
		return helpers.ApiResponse{}, err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := fmt.Errorf("unable to create request object: %s", err)
		return helpers.ApiResponse{}, err
	}

	request.Header = headers

	// do request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err := fmt.Errorf("unable to send request: %s", err)
		return helpers.ApiResponse{}, err
	}
	defer response.Body.Close()

	fmt.Println(" type of response in epic", reflect.TypeOf(response))

	// control status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return helpers.ApiResponse{}, fmt.Errorf("unexpected status code: %v", response.Status)
	}

	// read body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return helpers.ApiResponse{}, fmt.Errorf("error reading response body :%v", err)
	}

	fmt.Println("type of body", reflect.TypeOf(body))

	var result helpers.ApiResponse
	if err := utils.FromJSON(body, &result); err != nil {
		fmt.Println("Can not unmarshall JSON")
	}

	return result, nil
}
