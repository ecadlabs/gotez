package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ecadlabs/gotez/encoding"
)

type Client struct {
	Client *http.Client
	URL    string
	Chain  string
	APIKey string
}

type Error struct {
	Status int
	Raw    *http.Response
}

func newError(r *http.Response) *Error {
	return &Error{
		Status: r.StatusCode,
		Raw:    r,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("gotez: http status %d", e.Status)
}

func (c *Client) client() *http.Client {
	if c.Client != nil {
		return c.Client
	}
	return http.DefaultClient
}

func (client *Client) request(method string, path string, out any, ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, method, client.URL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/octet-stream")
	if client.APIKey != "" {
		req.Header.Add("X-Api-Key", client.APIKey)
	}
	res, err := client.client().Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		return newError(res)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	_, err = encoding.Decode(body, out, encoding.Dynamic())
	return err
}
