package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// chan<- 表示写入通道， <- 为指向符号

var wg sync.WaitGroup

func main() {
	ch := out()
	out := in(ch)

	// 这里能直接读取
	for v := range out {
		fmt.Println(v)
	}
	wg.Add(1)

	// 这里，如果使用协程，需要使用 wg wait 否则主协程就退出了
	go func() {
		defer wg.Done()

		fmt.Println("in main")

		// 还是可以从一个关闭的通道读取数据
		for v := range out {
			fmt.Println(v)
		}
	}()

	// 这里，需要 sleep，否则上面这个协程跑不完
	//time.Sleep(time.Second * 3)
	// 不用 sleep，使用 wait 来保证之前的协程都完成了
	wg.Wait()
}

// 生产者，返回一个只读的通道，因为它本身就是负责写入通道的
func out() <-chan int {
	// 这里，声明了一个双向通道，用来写，返回读
	out := make(chan int)

	// 注意：wg.Add 位置应该放在这里，也即开启的协程同级别，不能放在  go func 里
	//wg.Add(1)

	go func() {
		// 注意，这里必须 close，如果不 close，这个通道还是打开状态，阻塞，但不会有协程往里面写数据了
		// 但这里报两个错，所有在读这个通道的地方，会报错（48行）
		defer close(out)

		// 这行代码必须写在这里，表示这个协程结束了
		//defer wg.Done()

		rand.Seed(time.Now().UnixNano())
		out <- rand.Intn(100)
	}()

	// 也可以放在这里
	//wg.Add(1)

	// 双向通道可以被转换成单向通道返回
	return out
}

func in(in <-chan int) <-chan int {
	out := make(chan int)
	//wg.Add(1)

	// 这里，必须用一个新的 goroutine 封装，否则会ui deadlock
	go func() {
		// 注意，这里必须 close，如果不 close，这个通道还是打开状态，阻塞，但不会有协程往里面写数据了
		// 如果没有这行，15 行会报错 goroutine 1 [chan receive]:
		// 因为 15行在读取通道，拿不到数据了
		defer close(out)
		//defer wg.Done()

		for v := range in {
			fmt.Println("init", v)
			out <- v * v
		}
	}()

	return out
}

func test() {
	// 有缓冲，是异步发生的
	c := make(chan int8, 1)
	c <- 1
	fmt.Println(<-c)
}

func test1() {
	ch1 := make(chan int, 1)
	go func() {

		ch1 <- 3

		fmt.Println("sleep 1")
		time.Sleep(2 * time.Second)
		fmt.Println("sleep 2")
	}()

	//
	//for {
	//	select {
	//	case <-ch1:
	//		{
	//			fmt.Println("ch1 pop one")
	//			time.Sleep(1 * time.Second)
	//		}
	//	}
	//}

}
