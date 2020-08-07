package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {
	//fmt.Println("参数列表：", flag.Args())
	go func() {
		for {
			time.Sleep(time.Second * 3)
			log.Println(add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
	/*
		fun := func(n int) int {
			return n << 1
		}

		a := fun(123)
		fmt.Println(a)

		arr := []int{1, 3, 5, 7}
		index := sort.Search(len(arr), func(i int) bool {
			return arr[i] == 3
		})

		fmt.Println("find", index)

		arr2 := []int{1, 3, 5, 7, 8, 10, 20}
		index2 := sort.Search(len(arr2), func(i int) bool {
			return arr2[i] > 8
		})

		fmt.Println("find", index2)
	*/
}

type Worker struct {
	name  string
	score int
	next  *Worker
}

type Option struct {
}
