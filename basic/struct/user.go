package main


import (
    "fmt"
    "sync"
)


func main() {
    u := new(user)

    u.lock.Lock()
    u.name = "柳余"
    u.lock.Unlock()

    fmt.Println(u)
}

type user struct {
    name string
    age int
    lock sync.Mutex
}