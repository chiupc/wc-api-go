package net

import (
	"github.com/chiupc/wc-api-go/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
