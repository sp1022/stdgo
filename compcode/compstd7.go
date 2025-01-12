package compcode

import "fmt"

/*7 字典***********************************************************************************
1 字典是无序的，无法确定字段的返回顺序。本质上字典使用hash函数来实现内部的映射关系
  var 字典名 map[key数据类型]值数据类型
2 和切片一样，var定义字典不能直接使用，需要make初始化才能使用
3 字典和切片一样，作为函数参数时，底层共享内部结构，因为函数体内对字典的修改也会反应到函数的外部
*/
func Compstd7_func() {
	fmt.Println("********************字典基本使用**************************")
	var mymap map[string]string
	mymap = make(map[string]string) //使用make初始化
	//新增
	mymap["key1"] = "value1"
	mymap["key2"] = "value2"
	fmt.Println(mymap)
	//修改
	mymap["key1"] = "value2"
	fmt.Println(mymap)
	//删除
	delete(mymap, "key2") //使用函数delete删除
	fmt.Println(mymap)
	//检索
	fmt.Println(mymap["key1"])

	//短变量创建map
	mymap2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(mymap2)
}

/*字典和切片一样，作为函数参数时，底层共享内部结构，因为函数体内对字典的修改也会反应到函数的外部*/
func change2(m map[string]string) {
	m["name"] = "jack"
}
func Compstd7_func2() {
	fmt.Println("********************字典的增删改查**************************")
	mymap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	change2(mymap) //通过函数新增了一个name
	fmt.Println(mymap)

	//底层共享数据，通过mymap2删除，则mymap1也会修改
	mymap2 := mymap
	delete(mymap2, "name")
	fmt.Println(mymap)  //map[key1:value1,key2:value2]
	fmt.Println(mymap2) //map[key1:value1,key2:value2]

	v, e := mymap2["name"] //可以判断是否存在键name
	if e {
		fmt.Println(v)
	}
}
