package main

import "fmt"

/*
for range 本质是个语法糖：

for_temp := range
len_temp := len(for_temp)
for index_temp = 0; index_temp < len_temp; index_temp++ {
	value_temp = for_temp[index_temp]
	index = index_temp
	value = value_temp
	original body
}
*/
// 这里，可以看到，在循环开始前，有两个全局变量 index、value，表示当前迭代到的 index、当前 index 对应的 value
// 在 for .. range 这个逻辑中，开始已经确定好了长度 len_temp，所以即便一直 append，也不会计入
// 而因为 value 是全局变量，所以其内存地址不会变，只是在每次迭代中重新赋值而已，所以如果后续

func main() {
	testStruct2()
}

func testAddress() {
	// 迭代时，`index` 与 `value` 的地址是不变的
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Printf("Index: %d Value: %d Value-Addr: %X ElemAddr: %X \n", index, value, &value, &nums[index])
	}
}

type T struct {
	k string
}

func testStruct() {
	fmt.Println("-------- 不会修改值 --------")
	var arr [5]T
	for _, e := range arr {
		e.k = "xxx"
	}

	for _, e := range arr {
		fmt.Println(e.k)
	}

	fmt.Println("-------- 修改值 --------")

	var arr2 [5]T
	for i, _ := range arr2 {
		arr2[i].k = "foo"
	}

	for _, e := range arr2 {
		fmt.Println(e.k)
	}
}

func testStruct2() {
	arr := []T{
		T{"a"},
		T{"b"},
		T{"c"},
	}
	fmt.Println(arr)

	for i, t := range arr {
		fmt.Println(i, t.k)
		t.k = "1"
		fmt.Println(i, t.k)
	}

	fmt.Println(arr)
}

func testAppend() {
	// 不会死循环
	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println("i :", i)
		v = append(v, i)
	}

	// 这里能全部打印出来
	fmt.Println(v)
}

func testMap() {
	slice := []int{0, 1, 2, 3}
	dict := make(map[int]*int)
	for index, value := range slice {
		fmt.Println(&index, &value)
		// 这里，本质指向的都是同一个地址，也即 &value
		dict[index] = &value
	}

	fmt.Println("-------- map --------")
	for k, v := range dict {
		fmt.Printf("%d => %d\n", k, *v)
	}
}
