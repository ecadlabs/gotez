package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
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
	Body   []byte
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

func (client *Client) request(ctx context.Context, method string, path string, payload, out any, proto *core.Protocol) error {
	url := client.URL + path
	if client.Logger != nil {
		client.Logger.Printf("%s %s", method, url)
	}

	var (
		req *http.Request
		err error
	)
	if method == "POST" {
		var body bytes.Buffer
		if err = encoding.Encode(&body, payload, encoding.Dynamic()); err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, url, &body)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/octet-stream")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return err
		}
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
		e := &Error{
			Status: res.StatusCode,
			Raw:    res,
		}
		body, err := io.ReadAll(res.Body)
		if err == nil {
			e.Body = body
		}
		return e
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	encCtx := encoding.NewContext()
	if proto != nil {
		encCtx = encCtx.Set(core.ProtocolVersionCtxKey, *proto)
	}
	_, err = encoding.Decode(body, out, encoding.Ctx(encCtx), encoding.Dynamic())
	return err
}
