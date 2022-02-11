package web

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
}

func New(timeout int) Client {
	client := Client{
		&http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
	return client

}

func (c *Client) Get(url string, header map[string]string) (code int, body string, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = c.toHeader(header)
	resp, err := c.Do(req)
	if err != nil {
		return resp.StatusCode, "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(b), err
}

func (c *Client) toHeader(header map[string]string) http.Header {
	h := http.Header{}
	for k, v := range header {
		h[k] = []string{v}
	}
	return h
}
