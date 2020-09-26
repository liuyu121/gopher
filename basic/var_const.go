package main

import (
	"fmt"
	"reflect"
)

const x = 123

type Kind uint

const (
	invalid Kind = iota
	int3x
)

func main() {
	fmt.Println(invalid, int3x)

	y := 123
	z := "xx"

	t := reflect.TypeOf(x)
	t2 := reflect.TypeOf(y)

	fmt.Println("为什么数字能变成 int 呢 : ", t.Kind() == reflect.Int, t.Kind(), reflect.Int, reflect.TypeOf(reflect.Int), reflect.ValueOf(reflect.Int))
	fmt.Println(t, t.Name(), t.Kind(), t2, t2.Name())

	m := make(map[int]int)
	tm := reflect.TypeOf(m)
	fmt.Println(tm.Kind(), tm.Name())

	s := []int{1}
	ts := reflect.TypeOf(s)
	fmt.Println(ts.Kind(), ts.Name())

	t3 := reflect.ValueOf(x)
	t4 := reflect.ValueOf(y)
	t5 := reflect.ValueOf(z)
	t6 := reflect.ValueOf(&s).Elem()
	fmt.Println(t3.CanAddr(), t4.CanAddr(), t5.CanAddr(), t6.CanAddr(), s[0])
}