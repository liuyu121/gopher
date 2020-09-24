package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main2() {
	tag := `validate:"number,min=1,max=1000"`
	args := strings.Split(tag, ",")
	fmt.Println(args)
	fmt.Println(args[1:])

	var min, max int
	fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &min, &max)
	fmt.Printf("min=%v, max=%v\n", min, max)
}

func test_map(m map[int]int) {
	m[0] = 1
}

// 传引用和引用类型是有区别的， slice 是引用类型。
// 可以参见：https://blog.thinkeridea.com/201901/go/shen_ru_pou_xi_slice_he_array.html
// map 与 struct 的区别可见：https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
// 具体可参见：https://juejin.im/post/6844903762079776775
func test_slice(a []int) {
	// 注意，传参都是值传递，所以，这里 a 的地址改变了
	// 但是 array 指针、len、cap 都没改变，所以是个浅拷贝
	//fmt.Printf("2-1 a %v address is %p AND array is %v \n", a, a, unsafe.Pointer(&a))
	fmt.Printf("1-test_slice a real address is : %p\n", &a) // 这个是 slice 真实的地址，发生了一次值拷贝，所以地址变了
	fmt.Printf("1-test_slice a data-ptr address is : %p\n", a) // 这个是 slice 结构底层数组的地址，和外层 a 是一致的
	fmt.Printf("2-0 low level array: %p\n", (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	fmt.Println("2-1-0 low level array: ", *(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))

	a = append(a, 1)
	//fmt.Printf("2-2 a %v address is %p AND array is %v \n", a, a, unsafe.Pointer(&a))
	fmt.Printf("2-1 low level array: %p\n", (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	fmt.Println("2-1-1 low level array: ", *(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
}

func main() {
	m := make(map[int]int)
	a := make([]int, 0, 10)
	fmt.Printf("0-test_slice a real address is : %p\n", &a)
	fmt.Printf("0-test_slice a data-ptr address is : %p\n", a)

	//fmt.Printf("1-0 a %v address is %p\n", a, a)
	// 获取底层数组数据
	fmt.Printf("1-1 low level array: %p\n", (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	fmt.Println("1-2 low level array: ", *(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))

	test_map(m)
	test_slice(a)

	// 这是因为 a 的长度还是 0，虽然底层数组改变了
	fmt.Printf("3-1 low level array: %p\n", (*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	fmt.Println("3-2 low level array: ", *(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	//fmt.Println(*(*[10]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))))
	//fmt.Printf("3-0 a %v address is %p\n", a, a)

	fmt.Println(m, len(m)) // 1
	fmt.Println(a, len(a)) // 0
}
