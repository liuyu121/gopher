package main

import(
    "fmt"
    "time"
    "math/rand"
)

const (
    height = 15
    width = 80
)

type Universe [][]bool

func init() {
    rand.Seed(time.Now().Unix())
}

func NewUniverse() Universe {
    var u = make(Universe, width)
    for i:=1; i<width; i++ {
        for j:=1; j<height; j++ {
            u[i] = append(u[i], false)
        }
    }

    return u
}

// 检测一个细胞是否存活 上下左右细胞的情况
func (u Universe) alive(x, y int) bool {
}

func (u Universe) show() {
    for _, p := range u {
        for _, q := range p {
            str := " "
            if q {
                str = "*"
            }
            fmt.Printf(str)
        }
        fmt.Printf("\n")
    }
}

func (u Universe) send() {
    total := 0
    bingo := 0
    for i, g := range u {
        for j := range g {
            total++
            m := rand.Intn(4) + 1
            if m%4 == 0 {
                u[i][j] = true
                bingo++
            }
        }
    }

    fmt.Println(bingo, total)
    fmt.Println(float64(bingo)/float64(total))
}

func main() {
    u := NewUniverse()
    u.send()
    u.show()
}