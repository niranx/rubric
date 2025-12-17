package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"text/tabwriter"
)

// Header represents a single HTTP header with its name and values.
type Header struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// Response represents the HTTP response with URL, status, and headers.
type Response struct {
	URL     string   `json:"url"`
	Status  string   `json:"status"`
	Headers []Header `json:"headers"`
}

// FormatHeaders formats and writes HTTP headers to w according to the specified format.
// Supported formats are "plain", "table", and "json".
// It returns an error if formatting fails.
func FormatHeaders(w io.Writer, url string, resp *http.Response, format string) error {
	switch format {
	case "plain":
		return formatPlain(w, url, resp)
	case "table":
		return formatTable(w, url, resp)
	case "json":
		return formatJSON(w, url, resp)
	default:
		return nil
	}
}

// formatPlain formats headers in plain text format and writes them to w.
// Output includes the URL and status line, followed by headers in "Name: Value" format.
func formatPlain(w io.Writer, url string, resp *http.Response) error {
	fmt.Fprintf(w, "URL: %s\n", url)
	fmt.Fprintf(w, "Status: %s\n\n", resp.Status)

	headers := getSortedHeaders(resp.Header)
	for _, header := range headers {
		for _, value := range header.Values {
			fmt.Fprintf(w, "%s: %s\n", header.Name, value)
		}
	}

	return nil
}

// formatTable formats headers in aligned table format and writes them to w.
// The table has columns for header names and values with aligned spacing.
func formatTable(w io.Writer, url string, resp *http.Response) error {
	fmt.Fprintf(w, "URL: %s\n", url)
	fmt.Fprintf(w, "Status: %s\n\n", resp.Status)

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintln(tw, "HEADER\tVALUE")
	fmt.Fprintln(tw, "------\t-----")

	headers := getSortedHeaders(resp.Header)
	for _, header := range headers {
		for _, value := range header.Values {
			fmt.Fprintf(tw, "%s\t%s\n", header.Name, value)
		}
	}

	return nil
}

// formatJSON formats headers as indented JSON and writes them to w.
// The output includes the URL, status, and headers in a structured JSON format.
func formatJSON(w io.Writer, url string, resp *http.Response) error {
	response := Response{
		URL:     url,
		Status:  resp.Status,
		Headers: getSortedHeaders(resp.Header),
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(response); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

// getSortedHeaders converts HTTP headers to a sorted slice of Header structs.
// Headers are sorted alphabetically by name for consistent output.
func getSortedHeaders(headers http.Header) []Header {
	names := make([]string, 0, len(headers))
	for name := range headers {
		names = append(names, name)
	}
	sort.Strings(names)

	result := make([]Header, 0, len(names))
	for _, name := range names {
		result = append(result, Header{
			Name:   name,
			Values: headers[name],
		})
	}

	return result
}
