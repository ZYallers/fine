package ferror

import (
	"gitlab.sys.hxsapp.net/hxs/fine/util/fcast"
	"net/http"
)

type Error struct {
	msg  string
	code int
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	var errStr string
	if e.code != 0 {
		errStr = fcast.ToString(e.code)
	}
	if e.msg != "" {
		errStr += ":" + e.msg
	}
	return errStr
}

func New(values ...interface{}) (err *Error) {
	err = &Error{code: http.StatusOK, msg: ""}
	for _, val := range values {
		switch v := val.(type) {
		case int:
			err.code = v
		case string:
			err.msg = v
		case error:
			err.msg = v.Error()
		}
	}
	return
}
