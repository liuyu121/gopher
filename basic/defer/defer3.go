package main

import "fmt"


func main() {
    fmt.Println(returnValues()) // 0
    fmt.Println(namedReturnValues()) // 1
}

func returnValues() int {
    var result int
    defer func() {
        result++
        fmt.Println("defer")
    }()
    return result
}

func namedReturnValues() (result int) {
    defer func() {
        result++
        fmt.Println("defer")
    }()
    return result
}

