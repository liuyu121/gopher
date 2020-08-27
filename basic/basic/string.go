package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "test 柳余"
	fmt.Println(str)
	fmt.Println("len(str)", len(str))
	fmt.Println("RuneCountInString", utf8.RuneCountInString(str))

	str2 := 'x' // print 120 rune
	fmt.Println(str2)
}
