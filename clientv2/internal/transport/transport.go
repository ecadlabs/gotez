package transport

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	client "github.com/ecadlabs/gotez/v2/clientv2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type Transport client.Client

func (c *Transport) client() *http.Client {
	if c.Client != nil {
		return c.Client
	}
	return http.DefaultClient
}

func (tr *Transport) mkURL(path string, params map[string]any) (*url.URL, error) {
	u, err := url.Parse(tr.URL)
	if err != nil {
		return nil, err
	}
	values := make(url.Values, len(params))
	for k, v := range params {
		switch x := v.(type) {
		case string:
			if x != "" {
				values[k] = []string{x}
			}
		case client.Flag:
			if x {
				values[k] = []string{"yes"}
			}
		default:
			tmp := reflect.ValueOf(v)
			if (tmp.Kind() != reflect.Interface && tmp.Kind() != reflect.Pointer && tmp.Kind() != reflect.Slice) || !tmp.IsNil() {
				values[k] = []string{fmt.Sprintf("%v", v)}
			}
		}
	}
	u.Path = path
	u.RawQuery = values.Encode()
	return u, nil
}

func (t *Transport) debug(format string, a ...any) {
	if t.DebugLogger != nil {
		t.DebugLogger.Printf(format, a...)
	}
}

func wrapErr(err error) error {
	return fmt.Errorf("gotez-client: %w", err)
}

func (tr *Transport) Request(ctx context.Context, method string, path string, params map[string]any, payload, out any) error {
	u, err := tr.mkURL(path, params)
	if err != nil {
		return wrapErr(err)
	}

	tr.debug("%s %s", method, u.String())

	var req *http.Request
	if method == "POST" {
		var body bytes.Buffer
		if err = encoding.Encode(&body, payload, encoding.Dynamic()); err != nil {
			return wrapErr(err)
		}
		req, err = http.NewRequestWithContext(ctx, method, u.String(), &body)
		if err != nil {
			return wrapErr(err)
		}
		req.Header.Set("Content-Type", "application/octet-stream")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return wrapErr(err)
		}
	}
	req.Header.Set("Accept", "application/octet-stream")
	if tr.APIKey != "" {
		req.Header.Add("X-Api-Key", tr.APIKey)
	}
	res, err := tr.client().Do(req)
	if err != nil {
		return wrapErr(err)
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		e := &client.Error{
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
		return wrapErr(err)
	}
	if _, err = encoding.Decode(body, out, encoding.Dynamic()); err != nil {
		return wrapErr(err)
	}
	return nil
}

func Stream[T any](ctx context.Context, tr *Transport, path string, params map[string]any) (<-chan *T, <-chan error, error) {
	u, err := tr.mkURL(path, params)
	if err != nil {
		return nil, nil, wrapErr(err)
	}
	tr.debug("GET %s", u.String())
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, nil, wrapErr(err)
	}
	req.Header.Set("Accept", "application/octet-stream")
	if tr.APIKey != "" {
		req.Header.Add("X-Api-Key", tr.APIKey)
	}
	res, err := tr.client().Do(req)
	if err != nil {
		return nil, nil, wrapErr(err)
	}
	if res.StatusCode/100 != 2 {
		e := &client.Error{
			Status: res.StatusCode,
			Raw:    res,
		}
		body, err := io.ReadAll(res.Body)
		if err == nil {
			e.Body = body
		}
		res.Body.Close()
		return nil, nil, e
	}
	streamCh := make(chan *T)
	errCh := make(chan error)
	go func() {
		defer func() {
			res.Body.Close()
			close(streamCh)
			close(errCh)
			tr.debug("gotez-client: stream closed")
		}()

		for {
			// decoder is zero copy so new buffer must be allocated for each read
			var l [4]uint8
			if _, err := io.ReadFull(res.Body, l[:]); err != nil {
				errCh <- wrapErr(err)
				return
			}
			len := binary.BigEndian.Uint32(l[:])
			buf := make([]byte, int(len))
			if _, err := io.ReadFull(res.Body, buf); err != nil {
				errCh <- wrapErr(err)
				return
			}
			v := new(T)
			if _, err = encoding.Decode(buf, v); err != nil {
				errCh <- wrapErr(err)
				return
			}
			// v refers buf
			streamCh <- v
		}
	}()

	return streamCh, errCh, nil
}
