package main

import (
    "fmt"
    "time"
)

func main() {
    stop := make(chan bool)
    go func(i int) {
        for {
            select {
            case <- stop:
                fmt.Println("stop goroutine : ")
                return
            default:
                fmt.Println("do some thing")
                time.Sleep(time.Duration(i) * time.Second)
                fmt.Println("awake...")
            }
        }
    }(1)

    time.Sleep(5 * time.Second)
    fmt.Println("stop-1")
    stop <- true
    fmt.Println("stop-2")
    time.Sleep(5 * time.Second)
}
