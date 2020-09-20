package algo

import (
	"fmt"
)

interface {}
type Boss struct {
	name()
}

type Employ struct {
	b Boss
}

func (b *Boss) name(interface{}) {
	fmt.Println("this boss ")
}

func (e *Employ) name() {
	fmt.Println("this employ ")
}

// Top K问题, 给定一个非空的整数数组，返回其中出现频率前K高的元素,  算法的时间复杂度必须优于 O(n log n) 
func TopK(nums []int, n int) []int {
	// map 维护词频
	dict := make(map[int]int)
	for _, v := range nums {
		if _, exists := dict[v]; exists {
			dict[v] = dict[v] + 1
		}
	}

	// 对该 dict 排序

	s := []int(nil)
	for k, v := range dict {
		// 使用 slice 维护 k 大数字
		if len(s) < n {
			s = append(s, k)
		} else {
			// 如果大于当前长度，如果当前 v 大于栈顶元素，写入
			if v > dict[s[0]] {
				s[0] = k
			}
		}
	}

	return s
}

// go 针对 map 排序，需要通过 struct 实现
type ns struct {
	k int,
	v int,
}

type myHeap []int

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

// 排序，s[0] 为最小值元素
func sortQ(nums []int, dict map[int]int) {
	h := new(myHeap)
	for i := len(nums); i > 0; i-- {
		h.Push(nums[i]) // all elements are the same
	}
	h.verify(t, 0)

	// 排序 nums

	// 遍历 map，
}
