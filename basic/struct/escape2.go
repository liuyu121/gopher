package main

import "fmt"

var a = new(struct{}) // 堆上

func main() {
	var b = new(struct{})
	fmt.Println(b/*这里单独打印了，所以 b escape 到堆上*/, a == b)
}