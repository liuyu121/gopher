package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	//"strconv"
	"time"
)

type Value interface{}

// @see https://juejin.im/post/5a6873fef265da3e317e55b6
// @see https://www.flysnow.org/2017/05/12/go-in-action-go-context.html

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 这里，需要使用一个新的协程来封装
	go func() {
		t2(ctx)
	}()

	err := t1(ctx)
	if err != nil {
		fmt.Println(err)
		cancel()

		// 这里 必须 sleep
		time.Sleep(1 * time.Second)
	}
}

// https://segmentfault.com/a/1190000022484275
func t1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("stop t1")
}

func t2(ctx context.Context) {
	select {
	case <-time.After(time.Millisecond * 500):
		fmt.Println("t2 done")
	case <-ctx.Done():
		fmt.Println("stop t2")
	}
}

func t() {
	i := 0
	for {
		i++
		fmt.Println("in t")
		ctx := context.Background()
		out := make(chan Value) // 双向
		go test(ctx, out)
		fmt.Println(<-out)
		time.Sleep(1 * time.Second)
		if i == 5 {
			ctx.Deadline()
		}
	}
	//ctx2 = context.WithDeadline(context.TODO(), time.)
}

// 这里的参数，如果写成 out chan Value(双向) 与 out chan<- Value（只写），很大区别
func test(ctx context.Context, out chan<- Value) {
	// 这是一个死循环，done 何时触发呢，退出的时候触发
	for {
		v, err := hello("x1")
		if err != nil {
			fmt.Println("-- in test error --")
			return
		}

		fmt.Println("-- in test --")
		select {
		case <-ctx.Done():
			fmt.Println("done")
		case out <- v:
			fmt.Println(" 。。。 send 。。。")
			//return nil3
		}
	}
}

func hello(value interface{}) (value2 interface{}, err error) {
	if value == nil {
		return "none", errors.New("xo")
	}

	return value, nil
}

func mo() {
	var key = "__trace_id__"
	ctx, cancel := context.WithCancel(context.Background())

	// 主要用来在 goroutine 中传递数据
	for i := 1; i < 3; i++ {
		var traceId = "id-" + strconv.Itoa(i)

		valueCtx := context.WithValue(ctx, key, traceId)

		go watch(valueCtx, key)
	}

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second) // 这里，如果不 sleep，主进程就结束了，无法观察到 stop 的输出
}

func watch(ctx context.Context, key interface{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop")
			return
		default:
			fmt.Println("__trace_id__ is", ctx.Value(key))
			time.Sleep(1 * time.Second)
		}
	}
}
