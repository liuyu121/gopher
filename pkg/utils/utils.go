package utils

import (
	"encoding/json"
)

// json，存在错误则返回空
func JsonOrEmpty(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(str)
}

