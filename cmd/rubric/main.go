package main

import (
	"fmt"
	"os"

	"github.com/niranx/rubric/internal/cli"
	httpClient "github.com/niranx/rubric/internal/http"
	"github.com/niranx/rubric/internal/output"
)

func main() {
	config, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	client := httpClient.DefaultClient()

	resp, err := client.FetchHeaders(config.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if err := output.Format(os.Stdout, config.URL, resp, config.Format); err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
		os.Exit(1)
	}
}
