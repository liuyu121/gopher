package main

// go 函数传参都是 传值，除了 slice、map、chan 这种，即便 struct 也是值传递，除非使用了 指针

import (
    "fmt"
)

type say interface {
    sayHi(v interface{})
}

type Node struct {
    name string
    key string
    value int
}

// 注意，返回 Node 与 *Node 的区别，前者复制后
func main() {
    node := newNode("zhangsan", 123)
    node3 := node

    fmt.Printf("node : %v %p\n", node, &node)
    fmt.Printf("node copty -> node3 %v %p\n", node3, &node3)
    // node.sayHi() // 不能调用

    /*
    node.name = "x1x2xe"
    fmt.Printf("node : %v %p\n", node, &node)
    fmt.Printf("node copty -> node3 %v %p\n", node3, &node3)
*/
    fmt.Println("----------")

    node2 := newNode2("zhangsan", 123)
    node4 := node2
    node2.sayHi(node2)

    fmt.Printf("node2 : %v %p\n", node2, &node2)
    fmt.Printf("node2 copty -> node4 %v %p\n", node4, &node4)

    /*
    node2.name = "liuduoyu"
    fmt.Printf("node2 : %v %p\n", node2, &node2)
    fmt.Printf("node2 copty -> node4 %v %p\n", node4, &node4)
    */
}

func newNode(key string, value int) Node {
    node := Node{"", key, value}
    node.name = "newNode"
    return node
}

// 这里 要看是 *Node 还是 Node 实现了 say 接口的 sayHi 方法
func newNode2(key string, value int) say {
    node := &Node{"", key, value}
    node.name = "newNode2"
    return node
}

func (node *Node) sayHi(v interface{}) {
    fmt.Println("hi i am node : ", v.(*Node).name)
}
