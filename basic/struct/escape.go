package main

import "fmt"

// 逃逸分析：https://xargin.com/addr-of-empty-struct-may-not-eq/
// 更多参考：https://zhuanlan.zhihu.com/p/58065429
func main() {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b)

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d, c == d) // fmt.Println 因为参数是 interface 类型，所以会发生逃逸分析
}
