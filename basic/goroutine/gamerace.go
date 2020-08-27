package main

import(
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {
    baton := make(chan int, 4)

    wg.Add(1)

    go race(baton)

    // 第 1 位，开始跑步
    baton <- 1

    wg.Wait()
}

func race(baton chan int) {
    var newID int

    // 接收到数据
    id := <-baton
    fmt.Printf("%d is running\n", id)

    if (id < 4) {
        newID = id + 1;
        fmt.Printf("%d is to the line\n", newID)
        // 这里发起一个 goroutine 但是如何工作的呢
        // 这里 应该也是在等待数据，该协程阻塞，拿不到不回去
        go race(baton)
    }

    // runing
    time.Sleep(100 * time.Millisecond)

    if id == 4 {
        fmt.Printf("%d finished, race done\n", id)

        wg.Done()
        close(baton)
        return
    }

    fmt.Printf("%d send to the %d\n", id, newID)
    baton <- newID
    
}