package utils

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func FetchHeaders(url string) (http.Header, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp.Header, nil
}

func PrintHeaders(headers http.Header) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	headerColor := color.New(color.FgCyan).SprintFunc()
	// valueColor := color.New(color.FgYellow).SprintFunc()

	for key, values := range headers {
		for _, value := range values {
			fmt.Fprintf(w, "%s:\t\t%s\n", headerColor(key), value)
		}
	}

}
