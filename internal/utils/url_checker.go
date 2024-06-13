package utils

import (
	"net/url"
)

func CheckURL(input string) (*url.URL, error) {
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}
