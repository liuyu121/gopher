package algo

import (
	"container/heap"
	"fmt"
	"testing"
)

// 定义一个 切片类型
type myHeap []int

// heap 包内部会使用 sort 包来进行排序，也即 heap.Interface 接口实现了 sort.Interface
// 下面的 Less、Swap、Len 方法是实现了 sort.Interface
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return 10
}

// 下面的 Pop、Push 是实现 heap.Interface 中 sort 之外的方法
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func Test(t *testing.T) {
	h := new(myHeap)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}

	fmt.Println(h)

	heap.Init(h)

	for i := 10; i > 0; i-- {
		heap.Push(h, i)
	}

	for i := 1; h.Len() > 0; i++ {
		x := Pop(h).(int)
		if i < 20 {
			Push(h, 20+i)
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}
