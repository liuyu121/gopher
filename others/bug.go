package main

import "fmt"

// 这里有 bug
// go tool compile -d=ssa/prove/debug=3 bug.go
func main() {
	rates := []int32{1, 2, 3, 4, 5, 6}
	for star, rate := range rates {
		if star+1 < 1 {
			panic("")
		}
		fmt.Println(star, rate)
	}
}
