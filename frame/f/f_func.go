package f

import (
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fapp"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/frouter"
)

func App() *fapp.App { return fapp.Now() }

func Router() frouter.Router { return fapp.Now().Router }
