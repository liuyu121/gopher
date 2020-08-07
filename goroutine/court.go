package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("liuyu", court)
	go player("kobe", court)

	// start
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 接收通道的数据
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s won\n", name)
			return
		}

		// 随机一个数
		num := rand.Intn(100)
		//fmt.Printf("got %d mod 13 %d\n", num, num%13)
		if num%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)

			// 我方随机到了该数，关闭通道
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++
		court <- ball
	}

}
