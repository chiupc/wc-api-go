package request

import (
	"io"
	"net/url"
)

// Request ...
type Request struct {
	Method   string
	Endpoint string
	Values   url.Values
	Body	 io.Reader
}
