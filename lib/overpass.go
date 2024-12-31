package lib

import (
	"encoding/json"
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"io"
	"net/http"
	"net/url"
)

func OverpassQuery[X any](query string) ([]X, error) {
	var overpassResp struct {
		Elements []X `json:"elements"`
	}

	// URL encode the query
	encodedQuery := url.QueryEscape(query)

	// Overpass API endpoint
	apiURL := fmt.Sprintf("https://overpass-api.de/api/interpreter?data=%s", encodedQuery)

	// Send HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	//var overpassResp OverpassWaysResponse
	err = json.Unmarshal(body, &overpassResp)
	if err != nil {
		return nil, err
	}

	// Check if way was found
	if len(overpassResp.Elements) == 0 {
		return nil, fmt.Errorf("elements not found")
	}

	return overpassResp.Elements, nil
}

func OverpassFile(fileContents []byte) ([]data.Element, error) {
	var overpassResp struct {
		Elements []data.Element `json:"elements"`
	}

	err := json.Unmarshal(fileContents, &overpassResp)
	if err != nil {
		return nil, err
	}

	// Check if way was found
	if len(overpassResp.Elements) == 0 {
		return nil, fmt.Errorf("elements not found")
	}

	return overpassResp.Elements, nil
}
