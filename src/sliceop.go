package main

import("fmt")

func main() {
    slice := []int{10, 20, 30, 40}

    // 迭代时，使用一个地址表示每次迭代的值，所以地址总是相同的
    for index, value := range slice {
        fmt.Printf("Index: %d Value: %d Value-Addr: %X ElemAddr: %X \n", index, value, &value, &slice[index])

    }
}