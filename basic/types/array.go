package main

import (
	"fmt"
)

func main() {
	init2Arr()
	init2Arr2()
}

func p() {
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println("a[1] = ", array[1])

	for i, x := range array {
		fmt.Printf("array[%d] = %d \n", i, x)
	}
}

func bar() {
	var arr [1e6]int

	foo(arr)
}

func foo(arr [1e6]int) {
}
