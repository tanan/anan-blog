package web

import (
	"bytes"
	"health-planet-api/handler"
	"io"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
}

func NewClient(timeout int) handler.Client {
	client := Client{
		&http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
	return &client

}

func (c *Client) Get(url string, header map[string]string) (code int, respBody []byte, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = c.toHeader(header)
	resp, err := c.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, b, err
}

func (c *Client) Post(url string, header map[string]string, body []byte) (code int, respBody []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return 0, nil, err
	}
	req.Header = c.toHeader(header)
	resp, err := c.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, b, err
}

func (c *Client) toHeader(header map[string]string) http.Header {
	h := http.Header{}
	for k, v := range header {
		h[k] = []string{v}
	}
	return h
}
