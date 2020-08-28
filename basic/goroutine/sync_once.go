package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func main() {
    n := 10

    for i:=1; i<=n; i++ {
        once.Do(testSync)
        fmt.Println("i : ", i)
    }
}

func testSync() {
    fmt.Println("~~ once ~~")
    return
}

