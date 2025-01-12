package funccode

import "fmt"

/*
1 匿名函数:匿名函数没有函数名字，可以直接赋值给变量，之后就可以多次调用
2 变长函数，如果不确定函数的传入值有几个，可以使用边长函数
3 回调函数：本质就是作为另外一个函数的参数，在函数体中，可以在适当的实际调用参数对应的函数，形成回调。主要应用在事件机制。
4 闭包,例子：使用闭包的累加器,函数执行完毕后就无法修改函数中变量的值；有了闭包以后，函数就是一个变量的值，只要这个值没有释放，
	就可以在后期进行修改
5 defer，用户对资源的释放
*/
func Print(s string) {
	fmt.Println("run ", s)
}

/*2变长函数，如果不确定函数的传入值有几个，可以使用边长函数*/
func Sum(ns ...int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return ret
}

/*
3 回调函数：本质就是作为另外一个函数的参数
*/
type OnSumBefore func(int) int
type OnSum func(int, int) int
type OnSumEnd func(string)

var SumBeforeEvent OnSumBefore
var SumEvent OnSum
var SumEndEvent OnSumEnd

func StartSum(a, b int, c string) int {
	t, f := 0, 0
	//判断释放的绑定事件，并按事件执行顺序执行
	if SumBeforeEvent != nil {
		t = SumBeforeEvent(a)
	}
	if SumEvent != nil {
		f = SumEvent(t, b)
	}
	if SumEndEvent != nil {
		SumEndEvent(c)
	}
	return f
}

//RegEvent注册事件实现
func RegEvent(f1 OnSumBefore, f2 OnSum, f3 OnSumEnd) {
	SumBeforeEvent = f1
	SumEvent = f2
	SumEndEvent = f3
}

/*
4 闭包,例子：使用闭包的累加器,把函数func(int) int当成一个整体返回
*/
func Adder(i int) func(int) int {
	ret := i
	return func(n int) int {
		ret += n
		return ret
	}
}
