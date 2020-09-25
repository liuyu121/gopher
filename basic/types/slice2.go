package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	tet := SliceInt2String2(s)
	fmt.Println(tet)
}

// 切片使用技巧：https://blog.thinkeridea.com/201901/go/slice_de_yi_xie_shi_yong_ji_qiao.html
func gtS() {
	s := make([]int, 100)
	s[20] = 100
	s1 := s[10:10]
	s2 := s1[10:20]

	fmt.Printf("%T\n", s) // []int 切片
	fmt.Println(reflect.TypeOf(s))

	ss := [...]int{1}
	fmt.Printf("%T\n", ss) // [1]int 数组
	fmt.Println(reflect.TypeOf(ss))

	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2) // 这里 s1 的 偏移 10，对于底层数组就是偏移 20，且不会发生越界

	// 下面一个例子
	s44 := make([]int, 5)
	s44 = []int{1, 2, 3, 4}
	fmt.Println(len(s44), cap(s44), s44)
	s44 = append(s44, 6)
	s44 = append(s44, 7)
	s44 = append(s44, 6)
	fmt.Println(len(s44), cap(s44), s44)
}

func SliceInt2String2(s []int) string {
	if len(s) < 1 {
		return ""
	}

	bb := make([]byte, 0, 255)
	bs := "hello"
	bb = append(bb, bs...) // 这里是展开 string 为 []byte
	fmt.Println(bs, bb)

	b := make([]byte, 0, 256)
	fmt.Println(b)
	// 这里 ... 是将 string 的 slice 展开成 bytes，因为 string 本质是 bytes 的切片
	// strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。
	fmt.Printf("valus is %v and type is %T and string() is: %v\n", s[0], strconv.Itoa(s[0]), string(49))

	b = append(b, strconv.Itoa(s[0])...)
	fmt.Println("b is:", b)

	//x := (strconv.Itoa(s[0])...)
	//fmt.Println("b2 is:", s[0], x)

	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}

	return string(b)
}
