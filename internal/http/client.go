package http

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 10 {
					return fmt.Errorf("stopped after 10 redirects")
				}

				return nil
			},
		},
	}
}

// FetchHeaders fetches the HTTP response from the given URL using a GET request.
// The caller is responsible for closing the response body.
// It returns the response or an error if the request fails.
func (c *Client) FetchHeaders(url string) (*http.Response, error) {
	resp, err := c.httpClient.Head(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}

func DefaultClient() *Client {
	return NewClient(10 * time.Second)
}
