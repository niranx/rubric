package main

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func fetchHeaders(url string) (http.Header, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return response.Header, nil

}

func main() {
	url := "https://gmail.com"

	headers, err := fetchHeaders(url)

	if err != nil {
		fmt.Println(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Fprintf(w, "Headers for %s:\n\n", cyan(url))

	for key, values := range headers {
		for _, value := range values {
			fmt.Fprintf(w, "%s:\t\t%s\n", cyan(key), value)
		}

	}
	w.Flush()

}
