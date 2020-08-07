package main

import (
    "fmt"
)


//type my []string

func main() {
    // 注意 长度为 0 和 >0，初始化的方式不一样，存在长度则会初始化为 零值
    /*
    slice := make([]string, 10, 10)
    fmt.Println(slice)
    fmt.Println(len(slice))

    slice = append(slice, "xx", "liuxuoyu")
    fmt.Println(slice)
    fmt.Println(len(slice))
    */

    var intn = make([]int, 0, 0)
    cap1 := cap(intn)
    i := 0
    max := 2000
    for {
        i++

        intn = append(intn, i)
        cap2 := cap(intn)
        if cap1 != cap2 {
            fmt.Println(cap1, cap2)
        }
        cap1 = cap2
        if i > max {
            break
        }
    }
}