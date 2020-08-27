package main

import "fmt"

type myslice []string

func (m myslice) echo() {
    fmt.Println("my slice is", m)
}

func main() {
    var slice []string

    fmt.Printf("slice is %T\n", slice)

    var m myslice
    m.echo()
    slice.echo() // 调用失败
}