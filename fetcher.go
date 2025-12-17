package main

import (
	"fmt"
	"net/http"
)

// FetchHeaders fetches the HTTP response from the given URL using a GET request.
// The caller is responsible for closing the response body.
// It returns the response or an error if the request fails.
func FetchHeaders(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}

	return resp, nil
}
