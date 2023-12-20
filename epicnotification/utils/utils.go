package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	helpers "epicnotification/epicnotification/helpers"
)

// *** CREATE HEADERS ***

func CreateHeaders() (http.Header, error) {
	/* CREATE HEADERS FOR SCRAPER FROM JSON */
	var headerMap map[string]string

	err := json.Unmarshal([]byte(helpers.RequestHeaders), &headerMap)
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
