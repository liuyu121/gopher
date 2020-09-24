package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type W struct {
	a byte  // 1
	b int32 // 4
	c int32 // 4
}

type W2 struct {
	a byte  // 1
	b int32 // 4
	c int64 // 8
}

// -NOTE：unsafe 包重要的能力
// 1、任何类型的指针和 unsafe.Pointer 可以相互转换。
// 2、uintptr 类型和 unsafe.Pointer 可以相互转换。

// -- @seehttps://juejin.im/post/6844904067244769294

// 先简单介绍下Golang指针
// 		*类型：普通指针，用于传递对象地址，不能进行指针运算。
//		unsafe.Pointer：通用指针类型，用于转换不同类型的指针，不能进行指针运算。
//		uintptr：用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。
//　　
// unsafe.Pointer 可以和 普通指针 进行相互转换。
// unsafe.Pointer 可以和 uintptr 进行相互转换。
//
// 也就是说 unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。
// uintptr这个类型，在golang中，字节长度也是与int一致。通常Pointer不能参与运算，比如你要在某个指针地址上加上一个偏移量，Pointer是不能做这个运算的，那么谁可以呢？就是uintptr类型了，只要将Pointer类型转换成uintptr类型，做完加减法后，转换成Pointer，通过*操作，取值，修改值，随意。

func main() {
	// unsafe sizeof
	var p float64 = 99
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p))) // uintptr
	fmt.Println(unsafe.Sizeof(p))                 // 9

	// struct 对齐，先以字段对齐，再以 word 对齐
	var w3 W // 这时结构体的值类型
	var w31 *W // 这是结构体的指针
	fmt.Println("w3 W：~~")

	// 注意，如果是结构体
	fmt.Println("结构体：", unsafe.Alignof(w3)) // 4  --> 也即字段内部是以 4 字节对齐
	fmt.Println("结构体：", unsafe.Alignof(w31)) // 8 本质是个指针啊

	// 不管是 w 还是以 *w，其结构内部字段的字节对齐都是一样的
	fmt.Println(unsafe.Alignof(w3.a)) // 1，byte 类型要填充3个字节
	fmt.Println(unsafe.Alignof(w3.b)) // 4
	fmt.Println(unsafe.Alignof(w3.c)) // 8

	fmt.Println(unsafe.Alignof(w31.a)) // 1，byte 类型要填充3个字节
	fmt.Println(unsafe.Alignof(w31.b)) // 4
	fmt.Println(unsafe.Alignof(w31.c)) // 8

	fmt.Println("~~")

	fmt.Println("w *W2：~~")
	var w *W2
	var w22 W2
	fmt.Println("结构体：", unsafe.Alignof(w22)) // 8 --> 结构体内部以 8 字节对齐
	fmt.Println("结构体：", unsafe.Alignof(w)) // 8 本质是个指针啊

	// 注意：这里只是获取 struct 的指针 *w 的大小，而不是其值的大小
	fmt.Println(unsafe.Sizeof(w)) // 4 or 8，本机是 8
	fmt.Println("*w size：", unsafe.Sizeof(*w))               // 16，因为这里有字节对齐 a+b、c
	fmt.Println("*w size：", unsafe.Sizeof(w22))               // 16，因为这里有字节对齐 a+b、c
	fmt.Println("reflect w : ", reflect.TypeOf(w).Align())   // 8 字节对齐的大小

	fmt.Println(unsafe.Alignof(w.a)) // 1，byte 类型要填充3个字节
	fmt.Println(unsafe.Alignof(w.b)) // 4
	fmt.Println(unsafe.Alignof(w.c)) // 8

	fmt.Println("~~")

	var w2 *W2
	// 要获取值类型的大小，需要对指针变量进行取值 64 位操作系统按照 8 字节对齐
	fmt.Println("reflect w2 : ", reflect.TypeOf(w2).Align()) // 8 字节对齐的大小

	fmt.Println(unsafe.Alignof(w2.a)) // 8
	fmt.Println(unsafe.Alignof(w2.b)) // 8
	fmt.Println(unsafe.Alignof(w2.c)) // 8

	//fmt.Println(reflect.TypeOf(w.a).FieldAlign()) // 8
}
