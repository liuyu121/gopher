package main

import "fmt"

type Demo struct {
	Age    int
	Name   string
	Gender bool
}
type DemoPtr *Demo

// 错误：https://stackoverflow.com/questions/44406077/golang-invalid-receiver-type-in-method-func
// Go语言中规定，只有类型（Type）和指向他们的指针（*Type）才是可能会出现在接收器声明里的两种接收器，
// 为了避免歧义，明确规定，如果一个类型名本身就是一个指针的话，是不允许出现在接收器中的。
// go 语言的那些坑：https://i6448038.github.io/2017/07/28/GolangDetails/
func (d *Demo) ModifyName(name string) {
	d.Name = name
}

func main() {
	d := Demo{
		123,
		"liuyu",
		true,
	}

	fmt.Println(d)
	d.ModifyName("xx")
	fmt.Println(d)
}