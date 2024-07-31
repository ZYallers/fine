package fclient

import (
	"net/http"
	"time"
)

const (
	ClientTimeout = 15 * time.Second
)

// HttpClient Customize http.Client to optimize the performance based on the original http.DefaultTransport
// @see https://www.loginradius.com/blog/async/tune-the-go-http-client-for-high-performance
var HttpClient *http.Client

func init() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns, transport.MaxConnsPerHost, transport.MaxIdleConnsPerHost = 100, 100, 100
	HttpClient = &http.Client{Transport: transport, Timeout: ClientTimeout}
}
