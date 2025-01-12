package compcode

import "fmt"

/*4 切片***********************************************************************************
  (1)为什么需要切片？
    由于数组作为函数的参数时是需要赋值副本，如果一个数组很大，那么数组的相关操作非常消耗资源，可以使用切片，
  切片作为参数时候，只需要复制一些标志头的信息，不需要复制底层的数据数据，效率更高。
  (2)切片可以动态扩展
  (3)切片的组成：指向底层数组的指针，长度，容量。切片的底层是数组结构
  (4)两种定义方法：
   var 切片名 []数据类型
   切片名 :=make([]数据类型，长度，容量),
   tips:一个切片的长度用len获取，容量用cap获取。
        容量可以省略(省略后长度=容量)，长度必须小于等于容量
		长度指的是当前元素的个数，容量值最大的存储空间的个数
  (5)切片复制的是标识头，而不是底层数组
*/
func Compstd6_func() {
	fmt.Println("********************切片基本使用***************************")
	//第一种定义var slince []string slince = make([]string, 6)
	var slince []string
	fmt.Println(len(slince)) //0
	fmt.Println(cap(slince)) //0
	//slice[0]="hell0" 下标越界，因为没有定义长度
	slince = make([]string, 6)
	slince[0] = "hello"
	slince[5] = "go"
	fmt.Printf("%#v\n", slince) //[]string{"hello","","","","","go"}

	//第二种定义
	slice2 := []int{1, 2, 3, 4, 5} //不能指定长度，否则为数组
	slice1 := slice2               //切片赋值，共享底层数据
	slice1[0] = 9                  //由于指向同一个底层数据，所以两个切片的值同时影响
	fmt.Println(slice1)
	fmt.Println(slice2)

	//第三种短变量定义
	slice3 := make([]int, 3, 10) //不赋值时候，默认为[0,0,0]
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	fmt.Println(slice3)
}

/*4.2 切片的切割***********************************************************************************
slice[i:j] 从slice下标为i(i从0开始)的元素开始切，切片长度为j-1
slice[i:j:k] 从sclice下标为i(i从0开始)的元素开始切，切片长度为j-1,容量为k-1
slice[i:]  从下标为i(i从0开始)的元素切到最后一个
slice[:j]  从slice下标为0的元素切到下标为j-1的元素，不包含j
slice[:]   从头切到尾，相当于拷贝
*/

func Compstd6_func2() {
	fmt.Println("********************切片切割***************************")
	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[:] //切片拷贝
	slice2[0] = 9
	slice2[4] = 7
	//共享底层数据
	fmt.Println(slice)
	fmt.Println(slice2) //[9 2 3 4 7]
	//切1
	s3 := slice2[3:]
	fmt.Println(s3)      //[4 7]
	fmt.Println(len(s3)) //2
	fmt.Println(cap(s3)) //2,因为切片是从下标为3开始切，所以容量为5-3=2

	//切2
	s3 = slice2[:3]
	fmt.Println(s3)      //[9 2 3]
	fmt.Println(len(s3)) //3
	fmt.Println(cap(s3)) //5,因为切片是从下标为0开始切，且没有k值([i:j:k])，所以容量为5

	//切3
	s3 = slice2[1:4]
	fmt.Println(s3)      //[2,3,4]
	fmt.Println(len(s3)) //3
	fmt.Println(cap(s3)) //4,因为切片是从下标为1开始切，所以容量为5-1=4

	//切4
	s3 = slice2[1:3:4]
	fmt.Println(s3)      //[2,3]
	fmt.Println(len(s3)) //2
	fmt.Println(cap(s3)) //3，因为切片指定了k值，所以容量为4-1=3
}

/*4.3 切片的扩容***********************************************************************************
通过append函数来实现，第一个参数为要扩展的切片，第二个为要追加的值
tips:
1 扩容，新的切片有足够的容量，则和源切片共享底层数据，即修改新的切片值，源端也会修改
2 扩容，新的切片没有足够的容量，则和源切片不共享底层数据(意味着重新生成底层数据)，新切片修改底层数据，源端不变
*/
func Compstd6_func3() {
	fmt.Println("********************切片扩容***************************")
	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[1:3]
	fmt.Println(len(slice2)) //2 [2 3]
	fmt.Println(cap(slice2)) //4,因为切片是从下标为1开始切，所以容量为5-1=4

	//slice2和slice3是两个切片,扩容
	slice3 := append(slice2, 11) //扩容，有足够的容量，slice slice2 slice3共享底层数据
	fmt.Println(len(slice3))     //3 [2 3 11]
	fmt.Println(cap(slice3))     //4
	//测试是否共享底层数据：测试说明共享底层数据
	slice3[0] = 77
	fmt.Println(slice)  //[1,77,3,4,5]
	fmt.Println(slice2) //[77,3]
	fmt.Println(slice3) //[77,3,11]

	//扩容，容量不够的情况,slice3不共享底层数据
	slice3 = append(slice2, 11, 12, 13) //slice2容量为4，现在扩容到5
	fmt.Println(len(slice3))            //长度为5
	fmt.Println(cap(slice3))            //容量为4*2=8
	fmt.Println(slice)                  //[1,77,3,11,5]
	fmt.Println(slice3)                 //[77,3,11 12 13]

	//测试容量不够后是否共享？发现扩容，容量不够情况下，不共享
	slice3[0] = 99
	fmt.Println(slice)  //[1,77,3,11,5]
	fmt.Println(slice3) //[99,3,11 12 13]

	slice2[0] = 100
	fmt.Println(slice)  //[1 100 3 11 5]
	fmt.Println(slice2) //[100 3]
	fmt.Println(slice3) //[99,3,11 12 13]
}

/*4.4 切片作为函数的参数***********************************************************************************
1 切片作为函数的参数，函数内部切片参数和外部切片参数共享底层数据，所以对函数内部切片修改影响外部的切片值
*/

func change(slice []int) {
	slice[1] = 99
}
func Compstd6_func4() {
	fmt.Println("********************切片作为函数参数***************************")
	slice_source := []int{1, 2, 3, 4}
	slice_source2 := []int{6, 7, 8}
	slice_source = append(slice_source, slice_source2...)
	fmt.Println(len(slice_source)) //9
	fmt.Println(cap(slice_source)) //10,测试发现切片扩容的容量总是偶数，如果slice_source>slince_souce2,则容量=slice_source*2
	change(slice_source)           //调用函数修改值，通过修改函数内部的值
	fmt.Println(slice_source)
}
