package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(12 / 5)
	fmt.Println(14 / 5)

	test1()
	test()
	s := "L fdph, L vdz, L frqtxhuhg"

	for i := 0; i < len(s); i++ {
		c := s[i]

		if (c >= 'a' && c <= 'z') || c >= 'A' && c <= 'Z' {
			switch c {
			case 'a':
				c = 'x'
			case 'b':
				c = 'y'
			case 'c':
				c = 'z'
			}
			c -= 3
		}
		fmt.Printf("%c", c)
	}

}

func test1() {
	var x int8 = 127
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x+1)
	fmt.Printf("%d\n", x+1)

	fmt.Printf("a %#v\n", 'a')
	fmt.Printf("z %#v\n", 'z')
	fmt.Printf("A %#v\n", 'A')
	fmt.Printf("Z %#v\n", 'Z')
	fmt.Printf("%#v\n", 97-65)

}

func test() {
	var cent = 0.1
	var x float64
	for i := 1; i <= 10; i++ {
		x += cent
	}

	fmt.Println(x)

	ob := 0.0
	for i := 0; i < 11; i++ {
		ob += 0.1
	}
	fmt.Println(ob)
	fmt.Println(ob - x)

	y := 0.0
	arr := [3]float64{0.05, 0.1, 0.25}
	for {
		n := rand.Intn(2)
		y += arr[n]

		fmt.Printf("$%5.2f\n", y)
		if y > 20 {
			break
		}

	}
}
