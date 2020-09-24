package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type ttt1 struct {
	str unsafe.Pointer
	len int
	cap int
}

func ttt1Of(sp *string) *ttt1 {
	return (*ttt1)(unsafe.Pointer(sp))
}

func main() {

	s2 := "hello"
	fmt.Println("len is : ", len(s2))
	// *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s2)) + uintptr(8)))
	// 拆解： uintptr(unsafe.Pointer(&s2)) 获取 s2 的地址 + uintptr(8) 也即一个 8 的数字，再转化成 *ArbitraryType 类型
	// 而 *ArbitraryType 类型可以与任何类型转换
	fmt.Println("len is : ", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s2)) + uintptr(8))))

	fmt.Println("~~~")

	n := 8
	pn := &n
	pgn := (*uintptr)(unsafe.Pointer(pn))
	fmt.Println(n)    // 8
	fmt.Println(pn)   // 0xc00001a0c0
	fmt.Println(pgn)  // 0xc00001a0c0
	fmt.Println(*pn)  // 8
	fmt.Println(*pgn) // 8
	fmt.Println(uintptr(16))
	fmt.Println("~~~")

	s := "2"

	p := ttt1Of(&s)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(p.str))                // unsafe.Pointer
	fmt.Println(reflect.TypeOf(p.len))                // int
	fmt.Println(reflect.TypeOf(p.cap))                // int
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))     // uintptr
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p.str))) // uintptr

	fmt.Println(reflect.TypeOf(p))     // *main.ttt1
	fmt.Println(reflect.TypeOf(p.str)) // unsafe.Pointer

	// 这里 可以转化成 byte 指针
	g := (*byte)(p.str)
	fmt.Println(p)
	fmt.Println(g)
	fmt.Println(*g)
	fmt.Println(*(*byte)(p.str))
	fmt.Println(reflect.TypeOf(g)) // *uintptr

	// 从 ttt111 的首地址往后， +8字节，就是 len 成员的地址了 转化成 uintptr 才可以进行计算
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println("len is : ", Len)
	fmt.Println(p.len)
	fmt.Println(p.cap)
}
