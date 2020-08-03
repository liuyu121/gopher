package main

import(
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

var (
    counter int64
    wg sync.WaitGroup
)

func main() {
    wg.Add(2)

    go incrCounter(1)
    go incrCounter(2)

    wg.Wait()

    fmt.Println("Final Counter:", counter)
}

func incrCounter(id int) {
    defer wg.Done()

    for count := 0; count < 2; count++ {
        // 原子性的 +1
        atomic.AddInt64(&counter, 1)

        //value := counter

        // 交还给 CPU 调度其他 goroutine
        runtime.Gosched()

        //value++
        //counter = value
    }
}