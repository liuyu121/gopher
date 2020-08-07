package utils

import (
	"bytes"
	"strings"
)

// 使用 + 号，类似的，还可以使用
// 1、fmt.Sprint(args)
// 2、string 底层本质是个 []byte，所有用 append，再转化成 string：string(append([]byte(str)))
func AppendSimple(args ...string) string {
	s := ""
	for _, str := range args {
		s += str
	}

	return s
}

// 使用 bytes 库的 Buffer, 允许多次写入 []byte 、string 、rune 等类型的数据，并一次性输出
func AppendWitherBuffer(args ...string) string {
	var s bytes.Buffer
	for _, str := range args {
		s.WriteString(str)
	}

	return s.String()
}

// 使用 builder     var buffer bytes.Buffer
func AppendWitherBuilder(args ...string) string {
	var s strings.Builder
	for _, str := range args {
		s.WriteString(str)
	}
	return s.String()
}


// 解析一个字符串，去掉多余的空格，然后返回一个以空格连接的字符串
//
func TrimAndJoinString(str string) string {
	arr := strings.Fields(strings.TrimSpace(str))
	return strings.Join(arr, " ")
}
