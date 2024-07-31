package genctrl

import (
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
)

type apiItem struct {
	Import     string
	FileName   string
	Module     string
	Version    string
	MethodName string
}

func (a apiItem) String() string {
	return fstr.Join([]string{a.Import, a.Module, a.Version, a.MethodName}, ",")
}
