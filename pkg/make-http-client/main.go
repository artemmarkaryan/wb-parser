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
	t.TLSHandshakeTimeout = time.Minute

	return &http.Client{
		Timeout:   10 * time.Minute,
		Transport: t,
	}
}
