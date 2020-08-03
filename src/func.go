package main

import(
    "fmt"
)

func main() {
    msg := "张三"
    var f = func() {
        fmt.Printf("msg is %s\n", msg)
    }
    f()
    msg = "李四"
    f()
}