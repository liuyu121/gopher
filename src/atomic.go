package main

import(
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

var (
    // shutdown 是通知正在执行的 goroutine 停止工作的标志
    shutdown int64

    // wg 用来等待程序结束
    wg sync.WaitGroup
)

func main() {
    wg.Add(2) // 2 个协程

    go doWork("A")
    go doWork("B")

    // 本协程执行 1s
    time.Sleep(1 * time.Second)

    // 停止工作
    fmt.Println("Shutdown Now")
    atomic.StoreInt64(&shutdown, 1)

    wg.Wait()
}

func doWork(name string) {
    defer wg.Done()

    for {
        fmt.Printf("Dloing %s Work\n", name)
        // 0.25 s
        time.Sleep(250 * time.Millisecond)

        // 检测 shutdown 信号
        num := atomic.LoadInt64(&shutdown)
        fmt.Println("num is : ", num)
        if atomic.LoadInt64(&shutdown) == 1 {
            fmt.Printf("Shutting %s Down \n", name)
            break;
        }
    }
}