package main

import(
    "fmt"
    "runtime"
    "sync"
)

var (
    counter int

    // wg 用来等待程序结束
    wg sync.WaitGroup

    // mutex 用来定义一段代码临界区
    mutex sync.Mutex
)

func main() {
    wg.Add(2)

    go incrCounter(1)
    go incrCounter(2)

    wg.Wait()
    fmt.Println("Final Counter: %d \\n", counter)
}

func incrCounter(id int) {
    defer wg.Done()

    for count := 0; count < 2; count++ {
        // mutex lock 整个代码区段
        fmt.Println("1. id is : ", id)
        mutex.Lock()
        {
            value := counter

            fmt.Println("2. id is : ", id)
            runtime.Gosched() // 即便这里退出线程了，调度器也能再次把 id=1 的 goroutine 调度进来
            fmt.Println("3. id is : ", id)

            value++

            counter = value
        }

        // 解锁
        mutex.Unlock()
    }
}