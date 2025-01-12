package compcode

import "fmt"

/*******************************************************************
1 type关键字
2 struct类型
3 struct嵌套
4 匿名struct类型
5 数组
6 切片
7 字典
8 range关键字
*/

/*1.1 使用type 新数据类型  基于的数据类型*********************************************/
type name string //定义新类型name

func Compfunc() {
	var myname name = "jack"
	myname2 := myname
	fmt.Printf("%T \n", myname2) //数据的类型为compcode.name
	fmt.Printf("%t \n", myname2 == "jack")
	//myname3 := "jack"
	//fmt.Printf("%t \n", myname2 == myname3) //这里错误，因为类型name(变量myname2)和string(变量myname3)不匹配，无法比较。如果想比较，看Compfunc2
}

/*1.2 使用type 别名=数据类型***************************************************/
type mystr = string

func Compfunc2() {
	var myname mystr = "jack"
	myname = "smith"
	fmt.Printf("%T \n", myname) //类型仍然为字符串类型,输出为string
	myname2 := "smith"
	fmt.Printf("%t \n", myname2 == myname) //true
}

/*1.3  go语言中内置的数据类型别名*******************************************/
func Compfunc3() {
	var a byte
	fmt.Printf("%T \n", a) //byte的别名为int8
	var b rune
	fmt.Printf("%T \n", b) //rune的别名为int32
}
