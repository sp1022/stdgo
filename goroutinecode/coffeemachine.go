package goroutinecode

import "sync"

//定义咖啡机
type CoffeeMachine struct {
	Name       string
	CoffeeName string
	Gopher                //获取使用权的职工
	Mlock      sync.Mutex //互斥锁
}
