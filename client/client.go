package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type Logger interface {
	Printf(format string, a ...any)
}

type Client struct {
	Client *http.Client
	URL    string
	APIKey string
	Logger Logger
}

type Error struct {
	Status int
	Raw    *http.Response
	Body   []byte
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

func (client *Client) request(ctx context.Context, method string, path string, params map[string]any, payload, out any, proto *core.Protocol) error {
	u, err := url.Parse(client.URL)
	if err != nil {
		return err
	}
	values := make(url.Values, len(params))
	for k, v := range params {
		switch x := v.(type) {
		case string:
			if x != "" {
				values[k] = []string{x}
			}
		case Flag:
			if x {
				values[k] = []string{"yes"}
			}
		default:
			values[k] = []string{fmt.Sprintf("%v", v)}
		}
	}
	u.Path = path
	u.RawQuery = values.Encode()

	if client.Logger != nil {
		client.Logger.Printf("%s %s", method, u.String())
	}

	var req *http.Request
	if method == "POST" {
		var body bytes.Buffer
		if err = encoding.Encode(&body, payload, encoding.Dynamic()); err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, u.String(), &body)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/octet-stream")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
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
