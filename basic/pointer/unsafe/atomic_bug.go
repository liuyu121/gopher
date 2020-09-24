package main

import (
	"sync/atomic"
)

type T3 struct {
	b int64
	c int32
	d int64
}

func main() {
	a := T3{}
	atomic.AddInt64(&a.d, 1)
}