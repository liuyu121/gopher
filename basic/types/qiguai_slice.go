package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50}
	slice2 := slice[1:3]
	fmt.Println(slice, slice2)

	s := []byte("")
	println(s) // 添加用于打印信息, println() print() 为go内置函数，直接输出到 stderr 无缓存

	s1 := append(s, 'a')
	s2 := append(s, 'b')

	// 如果有此行，打印的结果是 a b，否则打印的结果是b b
	//fmt.Println(s1, "===", s2)
	fmt.Println(string(s1), string(s2))
}
