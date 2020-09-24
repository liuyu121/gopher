package main

import (
	"fmt"
	"unsafe"
)

// 内存对齐参考文章：
// https://zhuanlan.zhihu.com/p/53413177
// 另一篇详细文章：https://ms2008.github.io/2019/08/01/golang-memory-alignment/
// Dig101:Go之聊聊struct的内存对齐：https://juejin.im/post/6844904067244769294

// 对齐规则
//		结构体的成员变量，第一个成员变量的偏移量为 0。往后的每个成员变量的对齐值必须为编译器默认对齐长度（#pragma pack(n)）或当前成员变量类型的长度（unsafe.Sizeof），取最小值作为当前类型的对齐值。其偏移量必须为对齐值的整数倍
//		结构体本身，对齐值必须为编译器默认对齐长度（#pragma pack(n)）或结构体的所有成员变量类型中的最大长度，取最大数的最小整数倍作为对齐值
// 结合以上两点，可得知若编译器默认对齐长度（#pragma pack(n)）超过结构体内成员变量的类型最大长度时，默认对齐长度是没有任何意义的
// 也即，需要先对齐成员变量，再对齐结构体本身


// sizeof：32，内存布局如下：Part1 内存布局：axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
type Part1 struct {
	a bool // 1 -> 4
	b int32 // 4->8
	c int8 // 8->16
	d int64 // 16-24
	e byte // 24 -> 32
}


// sizeof：24，内存布局如下：Part1 内存布局：abxx|cxxx|dddd|dddd|exxx|xxxx
type Part2 struct {
	a bool // 1 -> 2
	b int8 // 2 -> 4
	c int32 // 4 -> 8
	d int64
	e byte
}

func main() {
	so()
}

func s1() {
	str := "liuyu"

	// byte        alias for uint8
	var data []byte = []byte(str)
	fmt.Println(data)

	// rune 类型 rune        alias for int32
	fmt.Printf("string size: %d\n", unsafe.Sizeof('A')) // 4
	fmt.Printf("string size: %d\n", unsafe.Sizeof('1')) // 4

	// string 是结构体 src/runtime/string.go
	// string is the set of all strings of 8-bit bytes, conventionally but not
	// necessarily representing UTF-8-encoded text. A string may be empty, but
	// not nil. Values of string type are immutable.
	fmt.Printf("string size: %d\n", unsafe.Sizeof("A")) // 16
}


func so() {
	part1 := Part1{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))

	fmt.Println(unsafe.Alignof(part1.a)) // 1，byte 类型要填充3个字节
	fmt.Println(unsafe.Alignof(part1.b)) // 4
	fmt.Println(unsafe.Alignof(part1.c)) // 1
	fmt.Println(unsafe.Alignof(part1.d)) // 8
	fmt.Println(unsafe.Alignof(part1.e)) // 1

	part2 := Part2{}
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))


	fmt.Println(unsafe.Alignof(part2.a)) // 1，byte 类型要填充3个字节
	fmt.Println(unsafe.Alignof(part2.b)) // 1
	fmt.Println(unsafe.Alignof(part2.c)) // 4
	fmt.Println(unsafe.Alignof(part2.d)) // 8
	fmt.Println(unsafe.Alignof(part2.e)) // 1
}

