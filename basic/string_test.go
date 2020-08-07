package main

import "testing"

// https://zhuanlan.zhihu.com/p/82556475

/*
goos: darwin
goarch: amd64
BenchmarkTestSimple-8    	21976570	        48.8 ns/op
BenchmarkTestBuffer-8    	16817413	        69.5 ns/op
BenchmarkTestBuilder-8   	19515334	        59.0 ns/op
PASS
ok  	command-line-arguments	3.949s
*/

import (
	"testing"

	sss "gopher/pkg/utils"
)

var str1 = "hello"
var str2 = "world"

func BenchmarkTestSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sss.AppendSimple(str1, str2)
	}
}

func BenchmarkTestBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sss.AppendWitherBuffer(str1, str2)
	}
}

func BenchmarkTestBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sss.AppendWitherBuilder(str1, str2)
	}
}
