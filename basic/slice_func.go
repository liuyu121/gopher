package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{9, 10}, {1, 4}, {3, 6}, {8, 12}}

	fmt.Println(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	fmt.Println(nums)

	ans := [][]int{}
	cur := []int(nil) // 存储区间 [x, y] 表示 x-y 是当前重合的
	for _, v := range nums {
		if len(ans) == 0 || cur[1] < v[0] {
			// 如果当前为空集，或当前区间的末尾 小于 下一区间的开头，表示没有相交
			ans = append(ans, v)
			cur = v
		} else if cur[1] < v[1] {
			// 如果当前区间末尾 小于 下一区间末尾，表示有交集，右移
			cur[1] = v[1]
		}
	}
	fmt.Println(ans)

	fmt.Println("---")
	nums2 := [][]int{{9, 10}, {1, 4}, {3, 6}, {8, 12}}
	fmt.Println(merge(nums2))
}

func merge(intervals [][]int) [][]int {
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 合并
	res := [][]int{}
	cs := []int{}
	for _, s := range intervals {
		if len(res) == 0 || cs[1] < s[0] {
			res = append(res, s)
			cs = s
		} else if cs[1] < s[1] {
			cs[1] = s[1]
		}
	}
	return res
}
