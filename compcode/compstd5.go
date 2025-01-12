package compcode

import "fmt"

/*4 数组********************************************
  var 数组名 [数组长度]数组类型
  tips:
  (1)数组数据没有初始化赋值，默认每个元素都是0
  (2)数组之间可以赋值，前提是两个数组的类型和长度必须一致
  (3)数组的赋值或者作为函数的参数时，都是需要赋值副本，如果一个数组很大，那么数组的相关操作非常消耗资源
*/
func Compstd5_func() {
	var arrName [3]string
	fmt.Printf("%#v", arrName) //[3]string{"","",""} 没有赋值默认为“”
	fmt.Println()
	//通过下标进行赋值
	arrName[0] = "hello"
	arrName[1] = ","
	arrName[2] = "array"
	fmt.Println(arrName)

	//定义短变量
	arrname2 := [3]int{1, 2, 3}
	fmt.Println(arrname2)

	//可以用...,会自动根据元素个数确定数组长度
	arrname3 := [...]int{4, 5, 6}
	fmt.Println(arrname3)
}

//数组作为函数的参数调用，都是值传递
func Compstd5_func2(arr [10]int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		arr[i] = 2 * i
	}
}

//数组作为函数的参数调用，都是值传递
func Compstd5_func3(arr *[10]int) {
	for i, v := range *arr {
		arr[i] = v * 2
	}
}
