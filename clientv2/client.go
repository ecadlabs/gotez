// Package client is a very limited Tezos RPC client library
package client

import (
	"fmt"
	"net/http"
)

type Flag bool

type Logger interface {
	Printf(format string, a ...any)
}

type Client struct {
	Client      *http.Client
	URL         string
	APIKey      string
	DebugLogger Logger
}

type Error struct {
	Status int
	Raw    *http.Response
	Body   []byte
}

func (e *Error) Error() string {
	return fmt.Sprintf("gotez-client: http status %d", e.Status)
}
