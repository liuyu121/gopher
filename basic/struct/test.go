package main

import "fmt"

type response interface {
	Header()
	Body()
}

type u struct {
	Name  string
	Score int
}

// struct 可以嵌入 interface，但可以不实现、或实现某几个 接口的方法，
type ux struct {
	response
	u1 u
}

type XX interface {

}

type Set struct {
	items map[interface{}]struct{}
}
func News() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

func main() {
	var x XX = nil
	var y XX = nil
	fmt.Println(x == y)

	test2()
}

func test1() {
	// ux 实现了 interface，但是没实现任何一个函数
	uw := ux{}

	var r response = &uw

	fmt.Println(uw)
	fmt.Println(r)
}

func (r *ux) Header() {
	fmt.Println("in func Header()")
}

func test2() {
	// ux 实现了 interface，实现任何 Header 函数
	uw := ux{}

	u2 := u{
		"zhangsan", 100,
	}
	uw.u1 = u2

	var r response = &uw

	fmt.Println(uw)
	fmt.Println(r)

	r.Header()
	uw.Header()

	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x109d33e]
	// 因为没有实现 Body() 所以段错误
	//r.Body()
}
