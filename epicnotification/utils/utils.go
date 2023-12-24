package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// *** CREATE HEADERS ***

func CreateHeaders(requestHeaders string) (http.Header, error) {
	/* CREATE HEADERS FOR SCRAPER FROM JSON */
	var headerMap map[string]string

	err := json.Unmarshal([]byte(requestHeaders), &headerMap)
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
