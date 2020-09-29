package main

import (
	"fmt"
	"sync"
	"time"
)

var wgg sync.WaitGroup

// https://stackoverflow.com/questions/35036653/why-doesnt-this-golang-code-to-select-among-multiple-time-after-channels-work
// 因为 time.After 是一个函数，总是会返回一个 newTimer 的 chan，所以每次拿到都是新的
// 如果需要复用，则要使用变量保存 timeout := time.After(3 * time.Second)
func main() {
	wgg.Add(1)

	// 较短的定时器会一直执行：https://www.cnblogs.com/elaron/p/8716977.html
	// https://stackoverflow.com/questions/39212333/how-can-i-use-time-after-and-default-in-golang
	go func() {
		for {
			select {
			case <-time.After(time.Second * 10):
				fmt.Println("in 10s")
			case <-time.After(time.Second * 1):
				// 因为每次返回的都是新的，所以永远走这个分支？ why？
				// --> 因为每次都 select 了一次，所以更大时间的 timer 也不会被执行
				// select 每次都会计算 case 的，而每次计算的时候，1s 的时间总是会产生一个，也即这个 7s 相对这里不是绝对时间
				// 而是 case 拿到这个语句的时候的 「相对时间」
				// 直接看这个答案就行了：
				// https://stackoverflow.com/questions/34894927/golang-timeout-is-not-executed-with-channels
				fmt.Println("in 1s")
			}
		}

		wgg.Done()
	}()

	fmt.Println("wait....")
	wgg.Wait()
}
