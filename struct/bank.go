package main

import (
	"fmt"
	"sync"
)

// 定义一个银行对象,，直接使用 sync 包，可以提升 sync.Mutex 的方法至 Bank 级别
// 也即，这里是 sync.Mutex 的方法集提升
// 嵌入类型
type Bank struct {
	//sync.Mutex
	sync.RWMutex
	balance map[string]int // 余额
}

// 初始化新的银行
func NewBank() *Bank {
	return &Bank{
		balance: make(map[string]int),
	}
}

// 存钱 函数存在接收者则为方法
func (bank *Bank) Deposit(name string, amount int) {
	defer bank.Unlock()

	// 首先 lock
	bank.Lock()
	// 余额增加，下面这个 if 也没必要需要
	if _, ok := bank.balance[name]; !ok {
		bank.balance[name] = 0
	}
	bank.balance[name] += amount
}

func (bank *Bank) Withdraw(name string, amount int) int {
	defer bank.Unlock()

	// 先 lock
	bank.Lock()

	// 检测余额是否足够
	if _, ok := bank.balance[name]; !ok {
		return 0
	}

	balance := bank.balance[name]
	if balance < amount {
		amount = balance
	}

	bank.balance[name] -= amount
	return amount
}

func (bank *Bank) Query(name string) int {
	// 查询只需要加读锁
	defer bank.RUnlock()

	// 先 lock
	bank.RLock()

	if _, ok := bank.balance[name]; !ok {
		return 0
	}

	return bank.balance[name]
}

func main() {
	bank := NewBank()

	name := "张三"

	go bank.Deposit(name, 20)
	go bank.Withdraw(name, 120)
	go bank.Deposit(name, 1000)

	//time.Sleep(time.Second * 3)

	fmt.Println(bank.balance[name])
	fmt.Println(bank.Query(name))
}
