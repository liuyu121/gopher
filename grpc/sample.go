package main

import (
	"fmt"
)

func main() {
	a := 1
	a = 2
	if a > 2 {
		a += 1
	}

	fmt.Println(a)
}
