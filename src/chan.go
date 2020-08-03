package main

import(
   "fmt"
   "sync"
   "math/rand"
   "time"
)

var (
    wg sync.WaitGroup
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    // 无缓冲通道
    court := make(chan int)

    wg.Add(2)

    // 启动两个选手
    go player("柳余", court)
    go player("kobe", court)

    // 发球
    court <- 1

    // 等待游戏结束
    wg.Wait()
}

func player(name string, court chan int) {
    defer wg.Done()

    fmt.Println("go ", name)

    for {
        // 等待球被击打过来
        ball, ok := <-court
        if !ok {
            fmt.Printf("Player %s Won\n", name)
            return
        }

        // 选随机数 然后用这个数来判断我们是否丢球
        n := rand.Intn(100)
        if n % 13 == 0 {
            fmt.Printf("Player %s Missed\n", name)

            // close 关闭通道，表示输了
            close(court)
            return
        }

        fmt.Printf("Plaer %s Hit %d\n", name, ball)
        ball++

        // 将球打向对手
        court <- ball
    }
}