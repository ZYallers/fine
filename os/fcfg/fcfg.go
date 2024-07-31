package fcfg

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/spf13/viper"
)

func ReadConfig(file string, merge bool) error {
	if !ffile.Exists(file) {
		return fmt.Errorf("config file \"%s\" not exists", file)
	}
	viper.SetConfigFile(file)
	if merge {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	} else {
		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	}
	return nil
}

func GetEnv(key string) interface{}  { return getEnvObjValue(viper.Get(key)) }
func GetEnvString(key string) string { return getEnvStrValue(viper.GetString(key)) }

func Get(key string) interface{}                      { return viper.Get(key) }
func GetString(key string) string                     { return viper.GetString(key) }
func GetInt(key string) int                           { return viper.GetInt(key) }
func GetDuration(key string) time.Duration            { return viper.GetDuration(key) }
func GetInt64(key string) int64                       { return viper.GetInt64(key) }
func GetFloat64(key string) float64                   { return viper.GetFloat64(key) }
func GetBool(key string) bool                         { return viper.GetBool(key) }
func GetStringSlice(key string) []string              { return viper.GetStringSlice(key) }
func GetStringMap(key string) map[string]interface{}  { return viper.GetStringMap(key) }
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}
func IsSet(key string) bool               { return viper.IsSet(key) }
func InConfig(key string) bool            { return viper.InConfig(key) }
func AllKeys() []string                   { return viper.AllKeys() }
func AllSettings() map[string]interface{} { return viper.AllSettings() }

func getEnvStrValue(value string) string {
	// 检查值是否符合 "env{" + key + "}" 的格式
	if !strings.HasPrefix(value, "env{") || !strings.HasSuffix(value, "}") {
		return value
	}
	// 提取环境变量的键
	envKey := strings.TrimSuffix(strings.TrimPrefix(value, "env{"), "}")
	// 从环境变量中获取值
	realValue := os.Getenv(envKey)
	// 如果环境变量存在，返回其值，否则返回原始字符串
	if realValue != "" {
		return realValue
	}
	return value
}

func getEnvObjValue(value interface{}) interface{} {
	switch val := value.(type) {
	case string:
		return getEnvStrValue(val)
	case []string:
		for i, s := range val {
			val[i] = getEnvStrValue(s)
		}
		return val
	case []interface{}:
		for i, v := range val {
			val[i] = getEnvObjValue(v)
		}
		return val
	case map[string]interface{}:
		for k, v := range val {
			val[k] = getEnvObjValue(v)
		}
		return val
	case map[string]string:
		for k, v := range val {
			val[k] = getEnvStrValue(v)
		}
		return val
	case map[string][]string:
		for k, v := range val {
			for i, s := range v {
				v[i] = getEnvStrValue(s)
			}
			val[k] = v
		}
		return val
	}
	return value
}
