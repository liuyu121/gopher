package main

import (
	"fmt"
	"unsafe"
)

func testP(p *int) {
	x := 100
	*p = x
}

// 二级指针：指向指针的指针
func testP2(p **int) {
	x := 100
	fmt.Println("x address : ", &x)
	*p = &x
	fmt.Println("**p : ", **p)
}

func testslice() {
	s := []int{1, 2, 3}
	// unsafe.Pointer() 可以与任何指针转化
	// uintptr 可以进行指针计算
	fmt.Println("s len is : ", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8))))
}

func main() {
	testslice()
}

func test22() {
	var p int
	testP(&p)
	fmt.Println(p)

	var p2 *int
	fmt.Println("p address : ", p2)
	testP2(&p2)
	fmt.Println("p address : ", p2)
	fmt.Println(*p2)
}