package ftracing

import "gitlab.sys.hxsapp.net/hxs/fine/os/fgoid"

func Go(f func()) {
	go func(mainTraceId string) {
		defer func() { recover() }()
		goid := fgoid.Get()
		defer DelTraceID(goid)
		SetTraceID(goid, mainTraceId)
		f()
	}(GetTraceID(fgoid.Get()))
}
