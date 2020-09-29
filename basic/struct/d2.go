package main

import (
	"fmt"
	"unsafe"
)

type Demo2 struct {
	Age    int
	Name   string
	Gender bool
}

func (*Demo2) Print() {
	fmt.Print("test")
}

func main() {
	d := (*Demo2)(unsafe.Pointer(nil))
	fmt.Println("unsafe.Pointer(nil) address = ", unsafe.Pointer(nil))
	fmt.Println("d address = ", &d)
	fmt.Printf("d address =  %p\n", &d)
	d.Print() // test
}
