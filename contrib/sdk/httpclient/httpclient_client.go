package httpclient

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"time"

	"gitlab.sys.hxsapp.net/hxs/fine/errors/ferror"
	"gitlab.sys.hxsapp.net/hxs/fine/net/fclient"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"go.uber.org/zap"
)

// Client is a http client for SDK.
type Client struct {
	prefix  string
	request *fclient.Request
	logger  *zap.Logger
	Handler
}

func (c *Client) SetPrefix(prefix string) *Client {
	c.prefix = prefix
	return c
}

func (c *Client) SetContentType(contentType string) *Client {
	c.request.SetContentType(contentType)
	return c
}

func (c *Client) SetHeader(key, value string) *Client {
	c.request.SetHeader(key, value)
	return c
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	c.request.SetHeaders(headers)
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.request.SetTimeOut(timeout)
	return c
}

func (c *Client) Request(ctx *gin.Context, req, res interface{}) error {
	if err := c.prepareRequest(req); err != nil {
		if c.logger != nil {
			c.logger.Error("prepare request", zap.Error(err))
		}
		return ferror.New(http.StatusInternalServerError, err)
	}
	resp, err := c.request.Send()
	if err != nil {
		if c.logger != nil {
			c.logger.Error("request send", zap.Error(err))
		}
		return ferror.New(http.StatusInternalServerError, err)
	}
	if err := c.HandleResponse(resp, res); err != nil {
		if c.logger != nil {
			c.logger.Error("handle response", zap.Error(err))
		}
		return err
	}
	return nil
}

func (c *Client) prepareRequest(req interface{}) error {
	url, method, data, err := c.parseRequest(req)
	if err != nil {
		return err
	}
	if len(c.prefix) > 0 {
		url = c.prefix + `/` + fstr.TrimLeft(url, `/`)
	}
	if !fstr.ContainsI(url, httpProtocolName) {
		url = httpProtocolName + `://` + url
	}
	c.request.SetUrl(url)
	if fstr.Contains(fstr.ToUpper(method), http.MethodPost) {
		c.request.SetMethod(http.MethodPost)
		c.request.SetHeader(httpHeaderContentType, httpHeaderContentTypeForm)
		if len(data) > 0 {
			c.request.SetPostData(data)
		}
	} else if fstr.Contains(fstr.ToUpper(method), http.MethodGet) {
		c.request.SetMethod(http.MethodGet)
		if len(data) > 0 {
			queries := make(map[string]string, len(data))
			for k, v := range data {
				queries[k] = fmt.Sprint(v)
			}
			c.request.SetQueries(queries)
		}
	}
	return nil
}

func (c *Client) parseRequest(req interface{}) (path, method string, data map[string]interface{}, err error) {
	typeOfReq := reflect.TypeOf(req)
	if typeOfReq.Kind() != reflect.Ptr {
		err = fmt.Errorf("request must be a pointer")
		return
	}
	if typeOfReq.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("request must be a struct")
		return
	}
	if _, ok := typeOfReq.Elem().FieldByName("Meta"); !ok {
		err = fmt.Errorf("request must have meta field")
		return
	}
	data = make(map[string]interface{})
	valueOfEleReq := reflect.ValueOf(req).Elem()
	for i := 0; i < valueOfEleReq.NumField(); i++ {
		valueOfField := valueOfEleReq.Field(i)
		typeOfField := valueOfEleReq.Type().Field(i)
		if valueOfField.Kind() == reflect.Struct {
			if typeOfField.Name == "Meta" {
				path = typeOfField.Tag.Get("path")
				method = typeOfField.Tag.Get("method")
			} else {
				for j := 0; j < valueOfField.NumField(); j++ {
					typeOfSubField := valueOfField.Type().Field(j)
					if fieldFormTag := typeOfSubField.Tag.Get("form"); fieldFormTag != "" {
						split := strings.Split(fieldFormTag, ",")
						data[split[0]] = valueOfEleReq.Field(i).Field(j).Interface()
					}
				}
			}
		} else {
			if fieldFormTag := typeOfField.Tag.Get("form"); fieldFormTag != "" {
				split := strings.Split(fieldFormTag, ",")
				data[split[0]] = valueOfEleReq.Field(i).Interface()
			}
		}
	}
	return
}
