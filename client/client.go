package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ecadlabs/gotez/v2/encoding"
)

type Logger interface {
	Printf(format string, a ...any)
}

type Client struct {
	Client *http.Client
	URL    string
	Chain  string
	APIKey string
	Logger Logger
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
	url := client.URL + path
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return err
	}
	if client.Logger != nil {
		client.Logger.Printf("%s %s", method, url)
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
