package main

import(
    "fmt"
    "sync"
    "time"
)

type ServiceData struct {
	ch   chan int // 用来 同步的channel
	data []int    // 存储数据的slice
}

func (s *ServiceData) Schedule() {
	// 从 channel 接收数据
	for i := range s.ch {
        fmt.Println("append i:", i)
		s.data = append(s.data, i)
	}
}

func (s *ServiceData) Close() {
	// 最后关闭 channel
	close(s.ch)
}

func (s *ServiceData) AddData(v int) {
   fmt.Println("add data")
    s.ch <- v // 发送数据到 channel
}

func NewScheduleJob(size int, done func()) *ServiceData {
    s := &ServiceData{
        ch:   make(chan int, size),
        data: make([]int, 0),
    }

    fmt.Println("1-new")
    // 这个协程一直在等待数据
    go func() {
        // 并发地 append 数据到 slice
        fmt.Println("1-after new : ", done)

        // 在这里阻塞，直到 s.ch 收到数据
        s.Schedule()

        fmt.Println("2-after new : ", done)

        // 终止程序
        done()

        fmt.Println("3-after new : ", done)
    }()
    fmt.Println("2-new")

    return s
}

func main() {
    var (
        wg sync.WaitGroup
        n  = 10
    )
    c := make(chan struct{})

    // new 了这个 job 后，该 job 就开始准备从 channel 接收数据了
    s := NewScheduleJob(n, func() { c <- struct{}{} })

    println("~~~~")

    time.Sleep(time.Second * 3)
    wg.Add(n)
    for i := 0; i < n; i++ {
        go func(v int) {
            defer wg.Done()
            s.AddData(v)
        }(i)
    }

    // 等待所有协程完成
    wg.Wait()
    s.Close()
    // 这个 会阻塞在这里，直到通道没数据写入了
    // 因为 NewScheduleJob 的协程会往 c 里面写数据，直到所有的都写完，否则会一直阻塞
    <-c

    fmt.Println(len(s.data))
}

