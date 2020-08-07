package main

import (
	"fmt"
)


// 使用内置类型，定义新的类型 但不能隐式转换
type Durantion int64

// 定义一个接口
type notifier interface {
    notify()
    receive(u user)
}

type user struct {
	name  string
	email string
}

type admin struct {
    person user // 声明类型 不是嵌入类型
    user // 这才是嵌入类型
    level string
}

// 使用指针方法实现 notify 接口
func (u *user) notify() {
	fmt.Printf("Sending User Emmail To %s<%s>\n\n", u.name, u.email)
}

func (u *user) receive(from user) {
	fmt.Printf("Receiving From User Emmail %s<%s>\n\n", from.name, from.email)
}

func (u *user) changeEmail(email string) {
    u.email = email
}

func sendNotify(n notifier) {
    n.notify()
}
func receiveNotify(n notifier) {
    n.notify()
}

func init() {
    fmt.Println("In the init func")
}

func main() {
    admin := admin{
        user: user{
            name: "root",
            email: "yuliqun@163.com",
        },
        person: user{
            name: "yuliqun",
            email: "yuliqun@163.com",
        },
        level: "root",
    } 

    liuyu := user{"Liuyu", "liuyu@qq.com"}
    liuyu.notify()
    admin.notify() 
    admin.person.notify()

    sendNotify(&liuyu)
    liuyu.receive(admin.person)
    /*
    // 直接调用
    admin.notify() // 报错
    admin.person.notify()
*/
    // ./struct.go:25:9: cannot use int64(1000) (type int64) as type Durantion in assignment
//    var dur Durantion
 //   dur = int64(1000)

 /*
    liuyu := user{"Liuyu", "liuyu@qq.com"}
    liuyu.notify()

    // 这里是个指针
    zhangsan := &user{"ZhangSan", "zhangsan@qq.com"}
    zhangsan.notify()
    (*zhangsan).notify()

    // 调用指针类型的方法
    liuyu.changeEmail("liuduoyu@gmail.com")
    liuyu.notify()

    (&liuyu).changeEmail("xxooyu@gmail.com")
    liuyu.notify()

    zhangsan.changeEmail("zhangxx@gmail.com")
    zhangsan.notify()
    */
}
