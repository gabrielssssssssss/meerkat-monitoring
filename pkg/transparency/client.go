package transparency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
	url  string
}

func NewClient(url string) *Client {
	return &Client{
		http: &http.Client{},
		url:  url,
	}
}

func (c *Client) GetTreeSize() (*Tree, error) {
	url := c.url + "/get-sth"

	resp, err := c.http.Get(url)
	if err != nil {
		return nil, ErrExecRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody
	}

	var tree Tree
	err = json.Unmarshal(body, &tree)
	if err != nil {
		return nil, ErrDecodeJSON
	}

	return &tree, nil
}

func (c *Client) GetEntries(start, end string) (*Entries, error) {
	url := fmt.Sprintf("%s/get-entries?start=%s&end=%s", c.url, start, end)

	resp, err := c.http.Get(url)
	if err != nil {
		return nil, ErrExecRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody
	}

	var entries Entries
	err = json.Unmarshal(body, &entries)
	if err != nil {
		return nil, ErrDecodeJSON
	}

	return &entries, nil
}
