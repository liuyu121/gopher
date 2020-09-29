package main

import "fmt"

func main() {
	fmt.Println(f())  // 1
	fmt.Println(f2()) // 5
	fmt.Println(f4()) // 2
}

func f() (result int) {
	// 这里 直接使用了 result 的地址
	defer func() {
		result++
		fmt.Println("defer result = ", result) // 1
	}()

	// ret = result
	// result++ = 1
	// return
	return 0
}

func f2() (r int) {
	t := 5

	// 1. 赋值指令
	// r = t

	// 这里 t 是个临时参数，本质来说地址是不固定的
	defer func() {
		t = t + 5
		fmt.Println("defer t = ", t) // 10
	}()

	// ret = 5
	// defer：t = 10
	// return
	return t
}

// 返回 2
// 这个 r 在外层调用函数里已经分配好地址了
func f4() (r int) {
	// r = 1

	// 这里，传参形式进入
	// 2. 这里改的 r 是之前传值传进去的 r，不会改变要返回的那个 r 值
	// 如果是地址传入，也会修改值
	defer func(r *int) {
		*r++
		fmt.Println("defer r = ", *r)
	}(&r)

	// result = 1
	// defer  -> 计算的是 r 的值
	// return
	return 1
}
