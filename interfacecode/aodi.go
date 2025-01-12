package interfacecode

import "fmt"

//定义AoDi结构体实现Drive方法
type AoDi struct{}

//实现idrive.go的方法
//*AoDi叫做接收器
func (*AoDi) Drive(name string) {
	fmt.Println("Drive AoDi  " + name + " Car")
}
