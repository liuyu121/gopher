package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//  @see https://www.cnblogs.com/qcrao-2018/p/10964692.html

type cat struct {
	name string
}

type cat2 struct {
	name *string
}

func main() {
	// 不同的 cat
	name := "zhangsan"
	c := cat{name}
	c2 := cat2{&name}

	fmt.Println(unsafe.Sizeof(c)) // 16
	fmt.Println(unsafe.Sizeof(c2)) // 8

	fmt.Println("~~~")

	var t uintptr // uint64
	fmt.Println("uintptr sizeof : ", unsafe.Sizeof(t))

	mp := make(map[string]int)
	mp["liuyu"] = 148
	mp["zhangsan"] = 120

	count := **(**int)(unsafe.Pointer(&mp))

	fmt.Println(reflect.TypeOf(unsafe.Pointer(&mp))) // unsafe.Pointer
	fmt.Println(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}