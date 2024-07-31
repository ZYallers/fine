package fsafe

import (
	"fmt"
	"runtime/debug"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/os/flog"
	"go.uber.org/zap"
)

func Defer() {
	r := recover()
	if r == nil {
		return
	}
	var err error
	if e, ok := r.(error); ok {
		err = e
	} else {
		err = fmt.Errorf("%v", r)
	}
	stack := string(debug.Stack())
	flog.Use().Error(err.Error(), zap.String("stack", stack))
	msg := fmt.Sprintf("recover from panic:\n%s\n%s", err, stack)
	fmsg.Sender().Simple(msg, true)
}
