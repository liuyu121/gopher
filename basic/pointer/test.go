package main

import "fmt"

func testP(p *int) {
	x := 100
	*p = x
}

// 二级指针：指向指针的指针
func testP2(p **int) {
	x := 100
	fmt.Println("x address : ", &x)
	*p = &x
	fmt.Println("**p : ", **p)
}

func main() {
	var p int
	testP(&p)
	fmt.Println(p)

	var p2 *int
	fmt.Println("p address : ", p2)
	testP2(&p2)
	fmt.Println("p address : ", p2)
	fmt.Println(*p2)
}