package frouter

import (
	"errors"
	"reflect"
	"sync"
)

var regMtx sync.Mutex

func (r Router) Register(controllers ...interface{}) {
	regMtx.Lock()
	defer regMtx.Unlock()
	for _, controller := range controllers {
		if err := r.doRegister(controller); err != nil {
			panic(err)
		}
	}
}

func (r Router) doRegister(ctrl interface{}) error {
	valueOfCtrl := reflect.ValueOf(ctrl)
	if valueOfCtrl.Kind() != reflect.Ptr {
		return errors.New("controller must be a pointer")
	}
	if valueOfCtrl.IsNil() {
		return errors.New("controller is nil")
	}
	typeOfCtrl := valueOfCtrl.Type()
	for i := 0; i < typeOfCtrl.NumMethod(); i++ {
		method := typeOfCtrl.Method(i)
		if !method.IsExported() || method.Type.NumIn() != 3 {
			continue
		}
		typeOfReq := method.Type.In(2)
		if typeOfReq.Kind() != reflect.Ptr {
			continue
		}
		elemOfReq := typeOfReq.Elem()
		if elemOfReq.Kind() != reflect.Struct {
			continue
		}
		metaField, exist := elemOfReq.FieldByName(FieldMetaKey)
		if !exist {
			continue
		}
		route := Route{
			provider:    ctrl,
			handlerName: method.Name,
			handlerReq:  elemOfReq,
			handler:     valueOfCtrl.Method(i),

			Path:   metaField.Tag.Get(PathMetaKey),
			Method: metaField.Tag.Get(MethodMetaKey),
			Login:  metaField.Tag.Get(LoginMetaKey) == "true",
			Sign:   metaField.Tag.Get(SignMetaKey) == "true",
		}
		r[route.Path] = &route
	}
	return nil
}
