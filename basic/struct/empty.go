package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type n struct{}

type S struct {
	A struct{}
	B struct{}
}

type SX struct {
	n
	y int
}

var s S
var sx SX

func main() {
	x := n{}
	println(unsafe.Sizeof(x))  // 0
	println(unsafe.Alignof(x)) // 1

	println(unsafe.Sizeof(s)) // 0

	// 空结构体不占用空间
	println(unsafe.Sizeof(sx))  // 8
	println(unsafe.Alignof(sx)) // 8

	// 元素为空结构体的数组 也不占据空间 但是其长度还是 1000000000
	var xss [1000000000]struct{}
	println(unsafe.Sizeof(xss), len(xss), reflect.TypeOf(xss).Kind(), reflect.TypeOf(xss).Kind() == reflect.Array) // 0 1000000000 17 true

	// 使用 make 会确保完成全部内存分配和相关属性的初始化
	var x1 = make([]struct{}, 1000000000)
	println(unsafe.Sizeof(x1), unsafe.Alignof(x1)) // 24 = 8+8+8

	a := n{}
	b := n{}
	// 这里涉及到了逃逸分析，在没有 fmt 调用打印 a、b 时，分配在栈上，但是因为 println 打印，分配在堆上，所以地址一样
	println(&a, &b, &a == &b) // false ? true?
	println(&a)
	//fmt.Println(a, b)       // 这里直接这么打印 也会返回 true
	//fmt.Printf("a addr = %v, b addr = %v\n", a, b)
	//fmt.Printf("a addr = %p, b addr = %p\n", a, b)
	//fmt.Printf("a addr = %p, b addr = %p\n", &a, &b) // 这一句打开，会返回 true

	fmt.Println("~~~")

	//ix := 123
	//iy := ix
	//println(&ix, &iy, &ix == &iy) // false
}
