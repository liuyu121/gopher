package main

import "fmt"

func main() {
	fmt.Println(returnValues())      // 0
	fmt.Println(namedReturnValues()) // 1
}

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result

	// ret = 0
	// defer --> result = 1
	// return --> 0
}

func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result

	// ret = result = 0
	// defer result++ -> 1
	// return
}
