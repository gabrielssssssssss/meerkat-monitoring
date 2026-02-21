package transparency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) GetTreeSize(url string) (*Tree, error) {
	ctUrl := url + "ct/v1/get-sth"

	fmt.Println(ctUrl)
	resp, err := c.http.Get(ctUrl)
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

func (c *Client) GetEntries(url string, start, end int64) (*Entries, error) {
	ctUrl := url + fmt.Sprintf("ct/v1/get-entries?start=%v&end=%v", start, end)

	resp, err := c.http.Get(ctUrl)
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
