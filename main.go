package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://httpbin.org/get"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Println("\nHeaders:")
	fmt.Println("--------")

	for name, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

}
