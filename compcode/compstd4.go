package compcode

import "fmt"

/*4 匿名struct********************************************
   var 变量名 struct{
	   字段1 字段1类型
	   字段N 字段N类型
   }
   匿名类型没有type关键字，使用var关键字，一般用作全局的数据库配置文件。如果在函数外部定义需要var关键字，
   如果在函数内部定义，var关键字可以省略
*/
//匿名结构体,只能在main中使用
var config struct {
	uid string
	pwd string
}

func Compstd4_func() {
	//config.uid = "shipeng" //匿名只能在main函数中赋值使用，否则报错，间接证明匿名结构体为单实例
	//config.pwd = "pwd"
	//fmt.Printf("%T", config)
	//函数内部定义可以省略var
	config2 := struct {
		uid string
		pwd string
	}{"shipeng2", "pwd2"}
	fmt.Printf("%T", config2)
	fmt.Println()
}
