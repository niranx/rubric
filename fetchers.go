package main

import (
	"fmt"
	"net/http"
	"os"
)

// fetchHeaders makes an HTTP GET request to the provided URL and returns the response headers.
// It also prints the HTTP status code and exits the program if the request fails.
func fetchHeaders(url string) http.Header {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n\n", resp.Status)
	return resp.Header
}

// printHeaders displays the provided HTTP headers to stdout in a formatted list.
func printHeaders(headers http.Header) {
	fmt.Println("Headers:")
	fmt.Println("--------")

	for name, values := range headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}
