package main

import "fmt"

// for range 本质是个语法糖：
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }
// 这里，可以看到，在循环开始前，有两个全局变量 index、value，表示当前迭代到的 index、当前 index 对应的 value
// 在 for .. range 这个逻辑中，开始已经确定好了长度 len_temp，所以即便一直 append，也不会计入
// 而因为 value 是全局变量，所以其内存地址不会变，只是在每次迭代中重新赋值而已，所以如果后续

func main() {
	slice := []int{10, 20, 30, 40}

	// 迭代时，使用一个地址表示每次迭代的值，所以地址总是相同的
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d Value-Addr: %X ElemAddr: %X \n", index, value, &value, &slice[index])

	}
}

func testRange1() {
	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println("i :", i)
		v = append(v, i)
	}

	// 这里能全部打印出来
	fmt.Println(v)
}

func testRange2() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
	for index, value := range slice {
		fmt.Println(&index, &value)
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	for k, v := range myMap {
		fmt.Printf("%d => %d\n", k, *v)
	}
}
