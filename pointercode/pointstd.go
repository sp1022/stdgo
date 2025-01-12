package pointercode

import (
	"fmt"
	"time"
	"unsafe"

	. "github.com/stdgo/pointercode2"
)

/*
 1 指针基础使用 point1函数
var 指针名 指针数据类型
 指针名 := new(指针数据类型)
2 获取元素地址
3 unsafe包使用
4 指针的指针
5 指针的值传递和地址传递
*/
func Point1() {
	fmt.Println("*******************************01指针基础使用***********************************")
	var a int64 = 8
	var ptr *int64 = &a
	//通过指针修改值
	*ptr = 9
	fmt.Println("%v\n", a)

	var c float32 = 3.14
	//ptr:=&c  //重要：不能将c赋给ptr，因为int64和float不兼容
	ptr2 := &c
	fmt.Printf("%v\n", ptr2)
	fmt.Printf("%v\n", c)
	c = 6.28
	fmt.Printf("%v\n", *ptr2) //由于C值修改，并且ptr2指向c，所以*ptr2变成6.28

	//开辟内存，d是*int指针
	d := new(int)
	fmt.Printf("%v\n", *d)
	*d = 9
	fmt.Printf("%v\n", *d)
}

func Point2() {
	/*
	*******************************2 获取元素地址***********************************
	 */
	fmt.Println("*******************************02获取元素地址***********************************")
	slice := []int{3, 4, 5, 6, 7}

	for _, value := range slice {
		value = 10
		//每次修改迭代value的值，但是&value的地址却没有变化，说明value本质上不是切片的各个元素的值
		//而是一个临时变量
		fmt.Println(value)
		fmt.Println(&value)
	}
	//上面修改slice的值，输出任然是{3，4，5，6，7}而不是{10,10,10,10},说明上面value实际上是slice的拷贝
	fmt.Println(&slice)
	fmt.Println(slice)

}

func Point3() {
	/*
	*******************************3 unsafe包使用***********************************
	 */
	stu := new(Student)
	//stu.id = 2   跨包调用，id是私有变量，无法赋值
	//stu.name = "shipeng"  跨包调用，id是私有变量，无法赋值
	fmt.Printf("%v \n", stu)

	//使用unsafe是否可以呢？
	p := (*string)(unsafe.Pointer(stu)) //p指向stu
	*p = "jack"                         //突破第一个私有变量
	fmt.Printf("%v \n", stu)

	//第二个id需要指针进行运算，第一个是字符串长度为16->uintptr(16)
	//计算过程：unsafe.Pointer(stu)获取地址，然后转换为uintptr,
	ptr_id := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + uintptr(16)))
	*ptr_id = 1
	fmt.Printf("%+v \n", stu)
}

func Point4() {
	/*
	*******************************4 指针的指针***********************************
	 */
	var b int32 = 7
	p1 := &b
	pp := &p1
	**pp = 9
	fmt.Println(b)
	fmt.Println(*&b)
	fmt.Println(p1)
	fmt.Println(pp)
	fmt.Println(*p1)  //对p1的地址取值*P1=9
	fmt.Println(*pp)  //*pp取值是指向pp的内存地址
	fmt.Println(**pp) //**pp 9
	//fmt.Fprintln(*&*&b)
}

func Point5() {
	/*
	*******************************5 指针的值传递和地址传递***********************************
	 */
	arr := [1024]int{}
	for i := 1; i <= 1024; i++ {
		arr[i-1] = i
	}
	start := time.Now()
	sum := 0
	for i := 0; i < 2; i++ {
		change(arr)
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("change(arr)执行了10000000次，耗时：", elapsed)
	fmt.Println(arr) //计算1000000次后arr值仍然不变，这就是值传递

	start = time.Now()
	sum = 0
	for i := 0; i < 2; i++ {
		changeByAddress(&arr)
		sum++
	}
	elapsed = time.Since(start)
	fmt.Println("changeByaddress(&arr)执行了10000000次，耗时：", elapsed)
	fmt.Println(arr) //计算1000000次后arr值改变，这就是地址传递
	//并且地址传递的速度要更快
}

/*
*******************************5 指针的值传递和地址传递***********************************
 */
func change(arr [1024]int) {
	for i, v := range arr {
		arr[i] = v * 2
	}
}

func changeByAddress(arr *[1024]int) {
	for i, v := range *arr {
		arr[i] = v * 2
	}
}
