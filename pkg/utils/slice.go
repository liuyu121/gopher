package utils

// 从 slice 按照 index 删除元素 - 保持顺序
func removeIntSliceByIndex(a []int, index int) []int {
	if index >= len(a) {
		return a
	} else {
		return append(a[:index], a[index+1:]...)
	}
}


// 从 slice 按照 index 删除元素 - 不保持顺序
func removeIntSliceByIndex2(a []int, index int) []int {
	if index >= len(a) {
		return a
	} else {
		a[index] = a[len(a)-1]
		return a[:len(a)-1]
	}
}

// 从 slice 按照 element 删除元素
func removeIntSliceByElem(a []int, value int) []int {
	i := 0
	for _, val := range a {
		if val != value {
			a[i] = val
			i++
		}
	}

	return a
}