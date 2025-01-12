package compcode

import "fmt"

/*8 range关键字***********************************************************************************
1 使用range迭代切片和字典，数组
2 range循环创建了每个元素的副本，而不是直接返回对该元素的使用，所以在range修改元素的值，元素本身不会改变
*/
func Compstd8_func() {
	fmt.Println("*******************range使用*****************************")
	//循环数组
	arr := [3]int{1, 2, 3}
	for index, v := range arr {
		fmt.Printf("arr[%d]->%d", index, v)
		fmt.Println()
	}
	//循环切片
	slice := []int{4, 5, 6}
	for index, v := range slice {
		fmt.Printf("slice[%d]->%d", index, v)
		fmt.Println()
	}
	//循环字典
	mymap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for k, v := range mymap {
		fmt.Printf("map[%s]->%s", k, v)
		fmt.Println()
	}
}

//range
func Compstd8_func2() {
	slice := []int{1, 2, 3}
	for _, v := range slice {
		v = v * 2
	}
	fmt.Println(slice) //[1,2,3] 说明是值拷贝，slince本身不会改变

	//循环追加,更改了slice
	for _, v := range slice {
		slice = append(slice, v)
	}
	fmt.Println(slice) //[1,2,3,1,2,3]
	mymap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for k := range mymap {
		if k == "key1" {
			delete(mymap, k)
		}
	}
	fmt.Println(mymap)
}
