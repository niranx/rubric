package cli

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
)

// Config holds the parsed command-line configuration.
type Config struct {
	URL    string
	Format string
}

// ParseArgs parses command-line arguments and returns the configuration.
// It validates the URL format and output format, returning an error if either is invalid.
func ParseArgs(args []string) (*Config, error) {
	fs := flag.NewFlagSet("rubric", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	format := fs.String("format", "plain", "Output format (plain, table, json)")

	fs.Usage = func() {
		printUsage(fs.Output())
	}

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if fs.NArg() < 1 {
		return nil, fmt.Errorf("no URL provided")
	}

	if !isValidFormat(*format) {
		return nil, fmt.Errorf("invalid format: %s (must be plain, table, or json)", *format)
	}

	url := fs.Arg(0)

	if err := validateURL(url); err != nil {
		return nil, fmt.Errorf("invalid URL: %w\n", err)
	}

	return &Config{
		URL:    url,
		Format: *format,
	}, nil
}

// isValidFormat checks if the provided format string is one of the supported output formats.
func isValidFormat(format string) bool {
	validFormats := map[string]bool{
		"plain": true,
		"table": true,
		"json":  true,
	}

	return validFormats[format]
}

// validateURL checks if the provided URL is valid and uses http or https scheme.
// It returns an error if the URL is malformed, missing a scheme, using an unsupported scheme, or missing a host.
func validateURL(rawURL string) error {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("malformed URL: %w", err)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("URL must start with http:// or https://")
	}

	if parsedURL.Host == "" {
		return fmt.Errorf("URL must include a host")
	}

	return nil
}

// printUsage prints the usage message and examples to stderr.
func printUsage(w io.Writer) {
	usage := `Usage: rubric [OPTIONS] <URL>

Fetch and display HTTP response headers for the given URL.

Options:
  -format string
        Output format: plain, table, json (default "plain")

Arguments:
  URL    The URL to fetch headers from (must start with http:// or https://)

Examples:
  rubric https://www.google.com
  rubric --format table https://api.github.com
  rubric --format json https://example.com

`
	fmt.Fprint(w, usage)
}
