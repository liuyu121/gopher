package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 3
    taskLoad         = 1
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// 任务
	tasks := make(chan string, taskLoad)

	// workers
	wg.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go work(i, tasks)
	}

    fmt.Println(time.Duration(10))
    fmt.Println(time.Millisecond)
    fmt.Println(time.Duration(10) * time.Millisecond)
    time.Sleep(10000 * time.Millisecond)
	for j := 1; j <= taskLoad; j++ {
		tasks <- fmt.Sprintf("task %d", j)
	} 
	// 加到通道后，这里可以立即关闭，goroutine 还是可以从通道取到数据
	close(tasks)

	wg.Wait()
}

func work(id int, tasks chan string) {
	defer wg.Done()

    fmt.Printf("worker %d is starting, waiting for task\n", id)
    for range time.Tick(1000 * time.Millisecond) {
        fmt.Println("tick ...")
    }
    for {
        // 等待分配工作 阻塞在这里
        task, ok := <-tasks
        if !ok {
            // 通道空了，且被关闭了, worker 关闭
            fmt.Printf("task is done, shutdown worker %d\n", id)
            return
        }

        fmt.Printf("worker %d is working %s\n", id, task)

        n := rand.Int63n(100)
        time.Sleep(time.Duration(n) * time.Millisecond)

        fmt.Printf("worker %d completed %s\n", id, task)
    }
}
