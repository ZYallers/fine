package fphp

import (
	"strings"

	"github.com/leeqvip/gophp/serialize"
)

func Serialize(vars map[string]interface{}) (str string) {
	for k, v := range vars {
		sa, _ := serialize.Marshal(v)
		str += k + "|" + string(sa)
	}
	return
}

func Unserialize(str string) map[string]interface{} {
	vars := make(map[string]interface{}, 10)
	offset := 0
	strlen := Strlen(str)
	for offset < strlen {
		if index := strings.Index(Substr(str, uint(offset), -1), "|"); index < 0 {
			break
		}

		pos := Strpos(str, "|", offset)
		num := pos - offset

		varname := Substr(str, uint(offset), num)
		offset += num + 1
		data, _ := serialize.UnMarshal([]byte(Substr(str, uint(offset), -1)))
		vars[varname] = data

		bytes, _ := serialize.Marshal(data)
		offset += Strlen(string(bytes))
	}
	return vars
}
