package main

import (
	"fmt"
	"strings"
)

func main() {
	tag := `validate:"number,min=1,max=1000"`
	args := strings.Split(tag, ",")
	fmt.Println(args)
	fmt.Println(args[1:])

	var min, max int
	fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &min, &max)
	fmt.Printf("min=%v, max=%v\n", min, max)
}
