package main

import (
	"fmt"
	"time"
)

func main() {
	testchan()
}

func testchan() {
	nums := 10
	c := make(chan int, nums)
	defer close(c)

	// 这里必须
	go func() {
		fmt.Println("start goroutine...")
		for {
			n, ok := <-c; if !ok {
				break
			}
			fmt.Println(n)
		}
	}()

	time.Sleep(time.Millisecond * 100)

	for nums > 0 {
		fmt.Println("add : ", nums)
		c <- nums
		nums--
	}
	fmt.Println(len(c))

	time.Sleep(time.Millisecond * 300)
}

func testxx() {
	stop := make(chan bool)
	go func(i int) {
		for {
			select {
			case <-stop:
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
