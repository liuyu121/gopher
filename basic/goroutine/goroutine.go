package main

import(
    "fmt"
    "runtime"
    "sync"
    "time"
    "math/rand"
)

func main() {
    // 分配一个逻辑处理器给调度器使用
    var procNum = runtime.NumCPU()
    runtime.GOMAXPROCS(procNum)
    fmt.Printf("procs num %d \n", procNum)

    // wg 用来等待程序完成
    // WaitGroup 是一个计数信号量，如果其值大于 0，Wait 方法会阻塞。
    // 设置为 2，表示有两个 goroutine 正在运行
    var wg sync.WaitGroup
    // 计数器 = 2，表示有两个 goroutine
    wg.Add(2)

    fmt.Println("Start Goroutines")

    // 声明两个匿名函数，函数内部创建 goroutine
    go func() {
        // 在本匿名函数退出时，调用 Done 来通知 main 函数工作已完成
        defer wg.Done() // 延时执行 保证 Done 方法一定会被调用

        for count := 0; count < 3; count++ {
            for char := 'a'; char < 'a'+26; char++ {
                num := rand.Int31n(3)
                time.Sleep(time.Duration(num) * time.Second)

                fmt.Printf("char is %c sleep %d\n", char, num)
            }
        }
    }()

    go func() {
        defer wg.Done() // 延时执行

        for count := 0; count < 3; count++ {
            for char := 'A'; char < 'A'+26; char++ {
                num := rand.Int31n(3)
                time.Sleep(time.Duration(num) * time.Second)

                fmt.Printf("char is %c sleep %d\n", char, num)
            }
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
}