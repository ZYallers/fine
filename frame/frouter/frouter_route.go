package frouter

import "reflect"

type Route struct {
	Sign        bool
	Login       bool
	Path        string
	Method      string
	handlerName string
	handler     reflect.Value
	handlerReq  reflect.Type
	provider    interface{}
}

func (r *Route) Provider() interface{} {
	return r.provider
}

func (r *Route) Handler() reflect.Value {
	return r.handler
}

func (r *Route) HandlerName() interface{} {
	return r.handlerName
}

func (r *Route) HandlerReq() reflect.Type {
	return r.handlerReq
}
