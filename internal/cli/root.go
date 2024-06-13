package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/niranx/rubric/internal/utils"
	"github.com/spf13/cobra"
)

var (
	url string
)

var rootCmd = &cobra.Command{
	Use:   "rubric",
	Short: "header extractor",
	RunE:  fetch,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL to process")
}

func fetch(cmd *cobra.Command, args []string) error {
	if url == "" {
		return errors.New("no URL provided, please use the --url or -u flag to specify a URL")
	}
	if _, err := utils.CheckURL(url); err != nil {
		return err
	}

	headers, err := utils.FetchHeaders(url)
	if err != nil {
		return err
	}

	fmt.Printf("Headers for %s:\n", url)
	utils.PrintHeaders(headers)
	return nil
}
