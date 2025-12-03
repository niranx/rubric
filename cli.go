package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// parseArgs parses command line arguments, validates the provided URL, and returns it.
// It exits the program with a usage message if no URL is provided or if validation fails.
func parseArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	rawURL := os.Args[1]

	if err := validateURL(rawURL); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return rawURL
}

// printUsage displays the CLI usage and help information to stdout.
func printUsage() {
	fmt.Println("Rubric - HTTP Header Fetcher")
	fmt.Println("\nUsage:")
	fmt.Println("  rubric <url>")
	fmt.Println("\nExample:")
	fmt.Println("  rubric https://example.com")
}

// validateURL checks if the provided URL is valid.
// It ensures the URL has an http:// or https:// prefix and parses correctly.
func validateURL(rawURL string) error {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		return fmt.Errorf("URL must start with http:// or https://")
	}

	_, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	return nil
}
