package main

import (
	"net/http"
	"io"
)

func DefaultRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("Authorization", "Bearer " + Config.Token)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}

	return client.Do(req)
}
