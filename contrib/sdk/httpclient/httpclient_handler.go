package httpclient

import (
	"net/http"
	"reflect"

	"github.com/ZYallers/fine/encoding/fjson"
	"github.com/ZYallers/fine/errors/ferror"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/net/fclient"
	"go.uber.org/zap"
)

type Handler interface {
	HandleResponse(resp *fclient.Response, out interface{}) error
}

type DefaultHandler struct {
	Logger  *zap.Logger
	RawDump bool
}

func NewDefaultHandler(logger *zap.Logger, rawRump bool) *DefaultHandler {
	return &DefaultHandler{
		Logger:  logger,
		RawDump: rawRump,
	}
}

func (h DefaultHandler) HandleResponse(resp *fclient.Response, out interface{}) error {
	if h.RawDump && h.Logger != nil {
		h.Logger.Debug("raw dump", zap.String("raw", resp.Request.DumpAll()))
	}
	valueOfOut := reflect.ValueOf(out)
	if valueOfOut.Kind() != reflect.Ptr {
		return ferror.New(http.StatusInternalServerError, "response must be a pointer")
	}
	switch valueOfOut.Elem().Kind() {
	case reflect.String:
		valueOfOut.Elem().SetString(resp.Body)
		return nil
	case reflect.Struct:
		result := f.JsonResult{Data: out}
		if err := fjson.Unmarshal([]byte(resp.Body), &result); err != nil {
			return ferror.New(http.StatusInternalServerError, err.Error())
		}
		if result.Code != http.StatusOK {
			return ferror.New(result.Code, result.Msg)
		}
		return nil
	default:
		return ferror.New(http.StatusInternalServerError, "response must be a struct or string pointer")
	}
}
