package main

import "fmt"

func main() {
	a := []int64{1}
	s1 := make([]int64, 0)
	s1 = append(s1, a...)
	fmt.Println(s1) // [1]

	// [0 0 0 0 0 1]
	s1 = make([]int64, 5)
	fmt.Println(s1) // [0 0 0 0 0]
	s1 = append(s1, a...)
	fmt.Println(s1) // [0 0 0 0 0 1]

	s1 = make([]int64, 0, 5)
	fmt.Println(s1) // []
	s1 = append(s1, a...)
	fmt.Println(s1) // [1]

	s1 = make([]int64, 0, 5)
	s1[4] = 1 // panic: runtime error: index out of range [4] with length 0
	fmt.Println(s1)
}
