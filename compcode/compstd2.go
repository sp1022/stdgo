package compcode

import (
	"fmt"
	"unsafe"
)

/*2 struct类型********************************************
2.1 类型定义
2.2 struct长度
*/

/*2.1 类型定义*********************************************
type 类型名 struct{
	字段1 字段1类型
	字段2 字段2类型
	........
	字段N 字段N类型
}
*/
type Student struct {
	name   string
	age    int
	height float32
}

func Compstd2_func() {

	//2.1.1 第一种实例化方法
	var stu Student
	stu.age = 3
	stu.name = "jack"
	stu.height = 96.8
	fmt.Println(stu.name, stu.age, stu.height)

	//2.1.2 第二种方法
	stu2 := Student{}
	stu2.age = 3
	stu2.name = "lucky"
	stu2.height = 98.9
	fmt.Println(stu2.name, stu2.age, stu2.height)

	//2.1.3 第三种方法
	stu3 := Student{
		name:   "lucy",
		age:    3,
		height: 101.1, //这里的,不能少
	}
	fmt.Println(stu3.name, stu3.age, stu3.height)

	stu34 := &Student{
		name:   "lucy",
		age:    3,
		height: 101.1, //这里的,不能少
	}
	fmt.Println(stu34.name, stu34.age, stu34.height)

	//2.1.4 第四种方法
	stu4 := new(Student)
	stu4.name = "ma"
	stu4.age = 3
	stu4.height = 100
	fmt.Println(stu4.name, stu4.age, stu4.height)
}

/*2.2 struct长度***********************************************************
结构体student的长度占用是多少？
	type stu struct{
	xh string             ---->16字节
	name string           ---->16字节
	age int               ---->8字节
	height float32        ---->4字节，但是由于内存对齐，由于age占用了8个字节，所以heignt需要填充4个字节，所以占用8个字节
	class string          ---->16字节
	}
	总共占有字节16+16+8+4+4(对齐)+16=64字节
 重要（注:go语言使用struct作为数据载体，由于指针对齐的问题，因此不同字段顺序可能导致结构体的总体大小也有所不同）
*/
type stu struct {
	xh     string
	name   string
	age    int
	class  string
	height float32
}

func Compstd2_func2() {
	stuvar := stu{
		xh:     "00000",
		name:   "shipeng",
		age:    20,
		height: 1.7,
		class:  "3",
	}
	fmt.Printf("stu占用空间为: %v", unsafe.Sizeof(stuvar)) //使用println会把%v直接打印出来，需要用printf格式化
}
