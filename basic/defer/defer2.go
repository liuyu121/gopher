package main

import "fmt"

func main() {
	//a := 10
	b := 10
 
	defer func() {
 
	}()
	defer func(b int) {
		fmt.Println("defer2",b)
	}(b)
	defer func() {
		fmt.Println("defer3",b)
	}()
 
	b = 20
	fmt.Println("main")
}
/* output:
main
defer3 20
defer2 10
*/
