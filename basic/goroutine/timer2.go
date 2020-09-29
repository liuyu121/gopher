package main

import (
	"fmt"
	"time"
)

// 问题在于：time.After 每次都会返回一个新的 channel，而 select 监听的是绑定的 select?
func main() {
	count := 0
	timeout1 := time.After(7 * time.Second)
	timeout2 := time.After(5 * time.Second)
	for i := 0; i < 36; i++ {
		select {
		case <-timeout1:
			fmt.Println(count, "7 second")
		case <-timeout2:
			fmt.Println(count, "5 seconds")
		default:
			fmt.Println(count, "Just wait")
		}
		count++
		time.Sleep(1 * time.Second)
	}
}
