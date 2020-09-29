package main

import (
	"encoding/json"
	"fmt"
	"modules/port-obj"
)

func main() {
	portobj.P(1)

	nt := t{
		"zhangsan",
		23,
	}

	fmt.Println(nt)

	str, _ := json.Marshal(nt)
	println(string(str))
}

type t struct {
	Name  string `json:"name"`
	// 注意 这里 不能有空格  json: "sc
	Score int `json:"score"`
}
