package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 传引用和引用类型是有区别的， slice 是引用类型。

const y = 123


func test_map(m map[int]int) {
	fmt.Printf("函数内 修改前 map 地址: s=%p &s=%p\n", m, &m)
	m[0] = 1
	fmt.Printf("函数内 修改后 map 地址: s=%p &s=%p\n", m, &m)
}

// 传引用和引用类型是有区别的， slice 是引用类型。
func test_slice(s []int) {
	// 注意，传参都是值传递，所以，这里 a 的地址改变了
	// 但是 array 指针、len、cap 都没改变，所以是个浅拷贝
	fmt.Printf("函数内 修改前 slice 地址: s=%p &s=%p data.ptr=%p\n", s, &s, (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
	fmt.Printf("函数内 修改前 slice=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("函数内 修改后 slice 地址: s=%p &s=%p data.ptr=%p\n", s, &s, (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
	fmt.Printf("函数内 修改后 slice=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))
}

func main() {
	fmt.Printf("y address is %p\n", y)
	t2 := reflect.ValueOf(y)
	fmt.Println(t2.CanAddr())
	//fmt.Println(&y) // cannot take the address of y

	m := make(map[int]int)
	fmt.Printf("外层 map 地址: s=%p &s=%p\n", m, &m)
	t := reflect.TypeOf(m)
	fmt.Println(t.Kind(), t.Name())

	s := make([]int, 0, 10)
	fmt.Printf("外层 slice 地址: s=%p &s=%p data.ptr=%p\n", s, &s, (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))

	test_map(m)
	test_slice(s)

	fmt.Printf("外层 map=%v, len(m)=%v\n", m, len(m))
	fmt.Printf("外层 slice=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))

	// 这是因为 a 的长度还是 0，虽然底层数组改变了
	//fmt.Printf("3-1 low level array: %p\n", (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
	//fmt.Println("3-2 low level array: ", *(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
	//fmt.Println(*(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	//fmt.Printf("3-0 a %v address is %p\n", a, a)
}
