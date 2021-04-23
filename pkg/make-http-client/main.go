package makeHTTPClient

import (
	"net/http"
	"time"
)

func NewHTTPClient(maxConnections int) *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = maxConnections
	t.MaxConnsPerHost = maxConnections
	t.MaxIdleConnsPerHost = maxConnections

	return &http.Client{
		Timeout:   time.Minute,
		Transport: t,
	}
}
