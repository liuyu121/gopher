package main

import "fmt"

func main() {
	t1_1()
	fmt.Println("")
	t2_1()
	fmt.Println("")

	a := 1
	b := 2
	fmt.Printf("a = %v, b = %v\n", a, b)
	a, b = b, a
	fmt.Printf("a = %v, b = %v\n", a, b)
}

// 4 3 2 1 0
func t1_1() {
	for i := 0; i < 5; i++ {
		defer func(x int) {
			fmt.Print(x, " ")
		}(i)
	}
}

// 5 5 5 5 5
func t2_1() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}