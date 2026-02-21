package telegram

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		http: &http.Client{
			Timeout: 30,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}
