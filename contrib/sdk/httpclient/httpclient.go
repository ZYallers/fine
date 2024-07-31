// Package httpclient provides http client used for SDK.
package httpclient

import (
	"fmt"
	"gitlab.sys.hxsapp.net/hxs/fine/net/fclient"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
)

const (
	httpProtocolName          = `http`
	httpHeaderContentType     = `Content-Type`
	httpHeaderContentTypeForm = `application/x-www-form-urlencoded`
)

// New creates and returns a http client for SDK.
func New(config Config) (client *Client) {
	client = &Client{
		request: fclient.NewRequest(""),
		Handler: config.Handler,
		logger:  config.Logger,
	}
	if client.Handler == nil {
		client.Handler = NewDefaultHandler(config.Logger, config.RawDump)
	}
	if !fstr.HasPrefix(config.Host, httpProtocolName) {
		config.Host = fmt.Sprintf("%s://%s", httpProtocolName, config.Host)
	}
	client.SetPrefix(config.Host)
	if config.Timeout > 0 {
		client.SetTimeout(config.Timeout)
	}
	return client
}
