package main

import (
	"fmt"
	"time"
)

var res = make(map[int][]int)

func main() {
	var s = []int{1, 2, 4, 6, 8}
	var e = []int{3, 5, 7, 9, 10}

	res := make(map[int][]int)    // 存入当前所有的解
	output := make(map[int][]int) // 存入某个 etime 对应的所有可选的集合

	etimes := []int{}
	index := 0 // 表示当前轮次可选取的集合
	curE := -1
	for i := 0; i < len(e); i++ {
		// 获取当前时间
		if curE == -1 {
			curE = e[i]
		}

		if curE == e[i] {
			// 如果不存在当前最优解中，存入当前轮次的可选解中，并更新当前最优解
			output[index] = append(output[index], i)
		} else {
			// 结束时间变了，遍历上一轮选取的结果，检测当前场次能否加入到新的一轮
			for j := 0; j < index; j++ {
				for _, k := range output[j] {
					if s[k] > s[i] {

					}
				}
			}

			// 更新当前最新的 etime
			curE = e[i]
		}

		// 如果当前场次的结束时间，与上一场的结束时间相同，检测开始时间是否有重叠
		if curE == e[i] {
			// 轮询某一轮可能出现的场次
			if cur, ok := output[i]; ok {
				// 如果存在当前最优解中
			} else {
				// 如果不存在当前最优解中，存入
				output[i] = append(output[i], i)
			}
		}
		// 拿到 curE 结束的所有场次
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这是个奇葩, 必须是这个时间点, 据说是go诞生之日, 记忆方法:6 - 1 - 2 - 3 - 4 - 5
}

func update() {

}
