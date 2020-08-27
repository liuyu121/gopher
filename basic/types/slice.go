package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	newSlice := append(slice, 888)

	// cap 倍数增加
	newSlice[2] = 19999
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(newSlice)
	fmt.Println(len(newSlice), cap(newSlice))
}

/*
func initTypes() {
    slice := make([]string, 5) // 长度 = 容量 = 5

    slice2 := make([]int, 3, 5) // 长度 = 3, 容量 = 5

    slice4 := []string{"Red", "Blue", "Green", "Yellow", "Plink"}


    slice5 := []string{99: ""}

    arr := [3]int{10, 20, 30} // [] 中指定了值，则为数组
    slice6 := []int{10, 20, 30} // [] 中未指定值，则为切片

    var slice7 []int // nil 切片 不作任何初始化时 长度 = 容量 = 0，nil 指针
    slice88 := make([]int, 0) // 空切片，长度 = 容量 = 0，地址指针
    slice89 := []int{} // 空切片，长度 = 容量 = 0，地址指针

//    slice3 := make([]int, 7, 5) // 长度 = 7, 容量 = 5 error
}
*/

func ref() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	slice[3] = 9999
	fmt.Println(slice)
	fmt.Println("----")

	newSlice := slice[1:3]
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println("----")

	// 使用新的 index
	slice[1] = 1009
	fmt.Println(slice)
	newSlice[1] = 21009
	fmt.Println(newSlice)
	fmt.Println(slice)
	fmt.Println("----")

	newSlice2 := append(newSlice, 60)
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	// newSlice[3] = 1009
	// slice[i:j:k]
	// 其中 i 表示从 slice 的第几个元素开始切，j 控制切片的长度(j-i)，k 控制切片的容量(k-i)，如果没有给定 k，则表示切到底层数组的最尾部。
}

func newSlice() {
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	fmt.Println(source)

	slice := source[1:2:3]
	fmt.Println(slice)

	slice2 := source[1:2:2] // 容量 = 长度，则 append 后，会生成新的底层数组
	fmt.Println(slice2)

	slice = append(slice, "x1") // 这时候，容量用完了
	fmt.Println(source)
	fmt.Println(slice)
	slice = append(slice, "x2") // 下面开始，就是使用新的底层数组
	fmt.Println(source)
	fmt.Println(slice)
	slice = append(slice, "x3")
	fmt.Println(source)
	fmt.Println(slice)
	slice = append(slice, "x4")
	fmt.Println(source)
	fmt.Println(slice)
	slice = append(slice, "x5")
	fmt.Println(source)
	fmt.Println(slice)
	fmt.Println(slice2)
	/*
	   slice2 = append(slice2, "x1y2")
	   fmt.Println(source)
	   fmt.Println(slice)
	   fmt.Println(slice2)
	*/
}

func nilSlice() {
	//var slice []int
	var slice []int
	slice[1] = 20
	fmt.Println(slice)

	//slice := make([]int, 0)
	//slice := []int{}
}

// 初始化二维数组
func init2Arr() {
	res := [][]int(nil)
	fmt.Println(res)

	row, col := 3, 4
	res = make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
	}
	fmt.Println(res)
}

func init2Arr2() {
	res := [][]int(nil)
	fmt.Println(res)

	row, col := 3, 4
	for i := 0; i < row; i++ {
		res = append(res, make([]int, col))
	}
	fmt.Println(res)
}

func capAndLen() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
