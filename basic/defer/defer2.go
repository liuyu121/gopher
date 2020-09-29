package main

import "fmt"

func main() {
	//a := 10
	b := 10
 
	defer func() {
 
	}()


	// 这里已经接收到了 b
	defer func(b int) {
		fmt.Println("defer2",b)
	}(b)

	// 这个 b 是全局的 b
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
