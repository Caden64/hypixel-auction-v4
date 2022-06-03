package HypixelRequests

import (
	"net/http"
	"time"
)

// NewClient returns a new http client to send requests
func NewClient() *http.Client {
	tr := &http.Transport{

		TLSHandshakeTimeout:   5 * time.Minute,
		ResponseHeaderTimeout: 5 * time.Minute,
		IdleConnTimeout:       5 * time.Minute,
		MaxConnsPerHost:       0,
	}

	c := http.Client{
		Transport: tr,
	}

	return &c
}
