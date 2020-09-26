package main

import "reflect"

type data struct{}

func (data) string() string {
	return "x"
}

type node struct {
	// 匿名接口类型
	data interface {
		string() string
	}
}

func main() {
	// 注意 这里 interface{string() string} 是一个定义，而 data 这个类型实现了 interface{string() string} 接口，所以 t 可以赋值、比较等
	var t interface {
		string() string
	} = data{}

	n := node{
		data: t,
	}

	tx := data{}
	tt := reflect.TypeOf(tx)
	println("tx type is ", tt.Name(), tt.Kind(), reflect.Struct)
	println(n.data.string())

	c := make(chan struct{})
	go func() {
		println("in go")
		c <- struct {}{}
	}()

	<-c
	println("done")
}
