package main

import "fmt"


func main() {
    fmt.Println(f()) // 1
    fmt.Println(f2()) // 5
    fmt.Println(f3()) // 1
}

func f() (result int) {
    // 这里 直接使用了 result 的地址
    defer func() {
        result++
    }()
    return 0
}

func f2() (r int) {
     t := 5

      // 1. 赋值指令
     // r = t

     // 这里 t 是个临时参数，本质来说地址是不固定的
     defer func() {
       t = t + 5
     }()
     return t
}


// 返回 1
func f3() (r int) {
    // r = 1

    // 这里，传参形式进入
    // 2. 这里改的r是之前传值传进去的r，不会改变要返回的那个r值
    // 如果是地址传入，也会修改值
    defer func(r *int) {
          *r++
    }(&r)

    // return
    return 1
}
