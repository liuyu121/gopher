package main

import (
	"fmt"
	"strconv"
	"strings"
)

// := 只能用在 func 内部
// 见语法糖：https://rainbowmango.gitbook.io/go/chapter10/1-foreword/1.1-warmup
func fun1() {
	i := 0

	// 这里 重新声明了 j，所以可以正常工作
	// 也即：no new variables on left side of :=
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}

var res = [][]int(nil)

func main() {
	/*
		var res2 = [][]int(nil)
		res2 = append(res2, []int{1, 2, 3})
		res2 = append(res2, []int{4})
		res2 = append(res2, []int{5, 0})
		fmt.Println(res2)
		fmt.Println(res2[0 : len(res2)-1])
	*/
	zijie()
	fmt.Println(res)
}

func zijie() {
	nums := []int{1, 2, 3}
	binary(nums)
	//list := []int(nil)
	//dfs(nums, 0, list)
}

// 为什么可以这么算呢，
func dfs(nums []int, start int, list []int) {
	fmt.Println("\ninit-", start, list)
	for i := start; i < len(nums); i++ {
		// 这里，每次都把当前的加进来
		list = append(list, nums[i])
		fmt.Println("loop-start", i, list)

		// 这里，递归加上后面的元素
		dfs(nums, i+1, list)
		list = list[0 : len(list)-1]
		fmt.Println("loop-after", i, list)
	}

	fmt.Println("break", list, "\n")
	res = append(res, list)
}

// 这个二进制，总长度
// 这里的重点是在于，把  nums 从 0~n-1 每一位，
// 对于 0~total-1 中的任一数字，比如 y ，那个位置上是1，就表示数组对应索引-1处的元素包含在内
// 那么，1 对应这个数组的哪一位置的元素呢，1 = 0001，表示第一个元素
// https://www.cnblogs.com/wizarderror/p/11364958.html
func binary(nums []int) {
	// 每一位，0、1 表示取（1）或者不取（0）
	length := len(nums)

	// 这是全部的子集个数，
	total := 1 << length
	// 下面是开始遍历整个子集（这里，我们假设已经知道了所有子集）
	for i := 0; i < total; i++ { // 从0～2^n-1个状态
		tmp := []int(nil)
		for j := 0; j < length; j++ { // 遍历二进制的每一位
			fmt.Println("loop", i, strconv.FormatInt(int64(i), 2), j, 1<<j, strconv.FormatInt(1<<j, 2), i&(1<<j))
			if i&(1<<j) == 0 { // 判断二进制第j位是否存在，这里只需要关心 j 这个位置，因为 1<<j 这个操作后， j 位置的值为 1，那么当且仅当 i 这个位置也为1，才表示 j 这个元素存在，因为当且仅当 1&1 == 0
				//if ((i >> j) & 1) == 1 {
				tmp = append(tmp, nums[j])
			}
			fmt.Println("tmp = ", tmp)
		}
		fmt.Println("final-tmp = ", tmp)
		res = append(res, tmp)
	}
}
func split() {
	str := "rescue rude erode you popular guard burden save forest laptop modify field"
	slice := strings.Split(str, " ")
	for _, s := range slice {
		fmt.Println(s)
	}
}
