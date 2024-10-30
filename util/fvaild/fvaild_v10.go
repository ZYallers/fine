package fvaild

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/ZYallers/fine/text/fstr"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	currentV10   *v10
	singletonV10 sync.Once
)

type v10 struct {
	Validate   *validator.Validate
	Translator translator.Translator
}

func V10() *v10 {
	lazyInit()
	return currentV10
}

func lazyInit() {
	singletonV10.Do(func() {
		currentV10 = &v10{}
		currentV10.Validate = validator.New()
		currentV10.Validate.SetTagName("validate")
		chinese := zh.New()
		uni := translator.New(chinese)
		currentV10.Translator, _ = uni.GetTranslator("zh")
		_ = translations.RegisterDefaultTranslations(currentV10.Validate, currentV10.Translator)
	})
}

func (v *v10) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyInit()
		if err := v.Validate.Struct(obj); err != nil {
			return errors.New(v.processError(obj, err))
		}
	}
	return nil
}

func (v *v10) Engine() interface{} {
	v.lazyInit()
	return v.Validate
}

func (v *v10) lazyInit() {
	lazyInit()
}

func (v *v10) processError(obj interface{}, err error) string {
	switch val := err.(type) {
	case *validator.InvalidValidationError:
		return val.Error()
	case validator.ValidationErrors:
		for _, item := range val {
			if v.Translator != nil {
				typeOf := reflect.TypeOf(obj)
				if typeOf.Kind() == reflect.Ptr {
					typeOf = typeOf.Elem()
				}
				trans := item.Translate(v.Translator)
				itemField := item.Field()
				if field, ok := typeOf.FieldByName(itemField); ok {
					if val := field.Tag.Get("form"); val != "" {
						return strings.Replace(trans, itemField, "["+val+"]", 1)
					}
				}
				return strings.Replace(trans, itemField, "["+fstr.CaseSnake(itemField)+"]", 1)
			} else {
				return item.Error()
			}
		}
	}
	return err.Error()
}
