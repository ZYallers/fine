package fjson

import (
	"github.com/json-iterator/go"
)

var Config = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	Config.RegisterExtension(&CustomTimeExtension{})
}

func Marshal(v interface{}) ([]byte, error) {
	return Config.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return Config.Unmarshal(data, v)
}

func Valid(data []byte) bool {
	return Config.Valid(data)
}
