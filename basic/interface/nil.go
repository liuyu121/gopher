package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 可参考：https://blog.csdn.net/qq_26981997/article/details/52608081
type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() {
	fmt.Println("meow")
}

func GetNilTabbyCat() Cat {
	var myTabby *Tabby = nil
	return myTabby
}
func GetTabbyCat() Cat {
	var myTabby *Tabby = &Tabby{}
	return myTabby
}

type Gafield struct{}

func (*Gafield) Meow() { fmt.Println("Gafield meow") }
func GetNilGafieldCat() Cat {
	var myGafield *Gafield = nil
	return myGafield
}
func GetGafieldCat() Cat {
	var myGafield *Gafield = &Gafield{}
	return myGafield
}

type iface struct {
	itype  uintptr
	ivalue uintptr
}

func main() {
	x := GetNilTabbyCat()

	xc := Cat()

	// 这里不为 nil，因为 type 不为 nil
	// interface 核心为 <type, value>，当且进度 type = nil && value = nil 也即 (nil, nil) 时才为 nil
	fmt.Println(x == nil, x)                           // false <nil>
	// 这里  x 是个指针，也即
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x), reflect.TypeOf(x).Kind()) // *main.Tabby <nil> ptr

	var cat1 Cat = nil
	cat2 := GetNilTabbyCat() // 接口指针
	cat3 := GetTabbyCat() // 接口指针
	cat4 := GetNilGafieldCat() // 接口指针
	var cat5 * Tabby = &Tabby{} // 结构体指针

	d1 := (*iface)(unsafe.Pointer(&cat1))
	d2 := (*iface)(unsafe.Pointer(&cat2))
	d3 := (*iface)(unsafe.Pointer(&cat3))
	d4 := (*iface)(unsafe.Pointer(&cat4))
	d5 := (*iface)(unsafe.Pointer(&cat5))

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println(d5)

	// cat2==cat5?:false, cat3==cat5?true
	// 也即 cat3 = cat5，接口指针与结构体指针相等
	fmt.Printf("cat2==cat5?:%5v, cat3==cat5?%v \n", cat2==cat5,cat3==cat5)
}
