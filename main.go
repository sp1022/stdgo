package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	. "github.com/stdgo/compcode"
	. "github.com/stdgo/funccode"
	. "github.com/stdgo/goroutinecode"
	. "github.com/stdgo/interfacecode"
	. "github.com/stdgo/interfacecode2"
	. "github.com/stdgo/interfacefunccode"
	. "github.com/stdgo/objectcode"
	. "github.com/stdgo/pointercode"
	. "github.com/stdgo/structcode"
)

//匿名结构体
var config struct {
	uid string
	pwd string
}

func generate1(ch chan<- int) {
	time.Sleep(3 * time.Second)
	ch <- 7
}
func generate2(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 8
}

/*
 gomod 和 gopath 两个包管理方案，并且相互不兼容，
 在 gopath 查找包，按照 goroot 和多 gopath 目录下 src/xxx 依次查找。
 在 gomod 下查找包，解析 go.mod 文件查找包，mod 包名就是包的前缀， 里面的目录就后续路径了。
 在 gomod 模式下，查找包就不会去 gopath 查找，只是 gomod 包缓存在 gopath/pkg/mod 里面。
 所以如果出现package Projectname is not in GOROOT，请GO111MODULE为off：go env -w GO111MODULE=off
*/

/*func test*/
func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	go Producer(ch)
	Consumer(ch)
	/*******CHAN错误例子0********************************************/
	//总结1：信道关闭后，从信写入取数据就会发生panic
	/*******CHAN错误例子1********************************************/
	// ch1 := make(chan int)
	// ch1 <- 7
	// n := <-ch //以下(包含本行)的代码无法执行，进入阻塞状态，等待其他携程从信道读取数据。在信道未关闭的情况下，从信道读取超时会引起deadlock
	// fmt.Println(n)
	// var input string
	// fmt.Scanln(&input)
	//总结2：信道必须有写有读，不然写完成后会阻塞引起deadlock
	/*******CHAN错误例子2***********************************************/
	// ch1 := make(chan int)
	// go func() {
	// 	ch1 <- 7
	// }()
	// //fmt.Println(<-ch1) //阻塞直到读取到7
	// var input string
	// fmt.Scanln(&input) //阻塞
	//总结3:一个信道必须有写有读；无写的时候，要关闭信道，否则会造成异常
	/*******CHAN错误例子3***********************************************/
	// ch1 := make(chan int)
	// go func() {
	// 	ch1 <- 7
	// 	close(ch1)
	// }()
	// fmt.Println(<-ch1) //读取到7
	// fmt.Println(<-ch1) //信道未关闭，读取到错误的0
	// var input string
	// fmt.Scanln(&input)
	//总结4：当携程从没有数据且未关闭的信道读取数据，就会进入阻塞状态，超时后抛出deadlock错误
	/************CHAN错误例子4***********************/
	// ch1 := make(chan int)
	// go func() {
	// 	ch1 <- 7
	// 	close(ch1)
	// }()
	// for {
	// 	if n, ok := <-ch1; ok {
	// 		fmt.Println(n)
	// 	} else {
	// 		break
	// 	}
	// }
	// var input string
	// fmt.Scanln(&input)
	//总结5：为了避免读取到0的错误数据，使用for.
	/************CHAN例子6***********************/
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	for n := range ch1 {
		fmt.Println(n)
	}
	//总结6：使用for range时候，如果信道未关闭，则会阻塞并引发deadlock
	/************CHAN例子7***********************/
	//总结7：没有close的信道不能用for range，只能用for,但是用for,最好也要写退出。这时可以使用select
	ch_1 := make(chan int)
	go generate1(ch_1)
	ch_2 := make(chan int)
	go generate2(ch_2)
	for { //select为无线循环
		select {
		case n1 := <-ch_1:
			fmt.Println(n1)
		case n2 := <-ch_2:
			fmt.Println(n2)
		case <-time.After(time.Second):
			fmt.Println("time out")
			goto EXIT //break无法退出for循环
		}
	}
EXIT:
	/************CHAN例子8***********************/
	ch_3 := make(chan int, 2)
	ch_3 <- 7
	ch_3 <- 8
	//ch_3 <- 9 //进入阻塞
	//总结8：有缓冲的信道，如果有容量N，超过该容量，则进入阻塞
	//总结9：一个已经关闭的信道内部的缓冲队列可能不是空的，在没有接收完这些值的情况下就关闭信道，会导致信道对象永远不会被
	//垃圾回收，也就是出现携程泄露的情况
	/*******************************************************************************************
	1匿名函数:匿名函数没有函数名字，可以直接赋值给变量，之后就可以多次调用
	*/
	f := func(a int, b int) (int, int) {
		return a * b, a + b
	}
	s, m := f(2, 3)
	fmt.Printf("第一次调用：%v  %v\n", s, m)
	s2, m2 := f(5, 6)
	fmt.Printf("第二次调用：%v  %v\n", s2, m2)

	/*****************************************************************************************
	2变长函数，如果不确定函数的传入值有几个，可以使用边长函数
	*/
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(22, 33, 44))

	/******************************************************************************************
	  3回调函数：本质就是作为另外一个函数的参数
	*/
	//定义三个函数
	f1 := func(a int) int {
		println("=======SumBeforeEvent==============")
		return a + 1
	}
	f2 := func(a, b int) int {
		println("==========SumEvent===============")
		return a + b
	}
	f3 := func(c string) {
		println("===========SumEndEvent===========")
		println(c)
	}
	RegEvent(f1, f2, f3)
	ff := StartSum(3, 7, "End")
	println(ff)

	/********************************************************************************************************
	4 闭包,例子：使用闭包的累加器,函数执行完毕后就无法修改函数中变量的值；有了闭包以后，函数就是一个变量的值，只要这个值没有释放，
	就可以在后期进行修改
	*/
	fc := Adder(2)
	fc(1) //2+1
	fc(2) //2+1+2
	fc(3) //2+1+2+3
	//8+5
	fmt.Println(fc(5))

	/**********************************************************************************************************
	5 defer
	*/
	fmt.Println("==========start=========")
	//defer执行顺序和调用顺相反,defer是从下往上返回
	defer print("order 1")
	defer print("order 2")
	defer print("order 3")
	fmt.Println("============end===========")

	/**********************************************************************************************************
	goroutine学习
	*/
	//模拟双核
	runtime.GOMAXPROCS(2)
	//wg.add用来阻塞等待2个协程任务完成
	var wg sync.WaitGroup
	wg.Add(2)

	gopher := Gopher{Name: "Gopher1", ID: 1}
	gopher2 := Gopher{Name: "Gopher2", ID: 2}

	t := time.Now()

	//gopher1和gopher2并发执行
	go gopher.MakeCoffee("A", &wg)
	go gopher2.MakeCoffee("B", &wg)
	//wg.wait用于阻塞主线程，并等待2个子协程完成
	wg.Wait()

	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
	fmt.Println("================END===================")
	/***************************************************************************************
	面向对象
	*/
	fmt.Println("***********************面向对象****************************")
	p := Personobj{}
	//p.name p.age私有变量无法访问，通过set方法
	p.SetName("jack")
	p.SetAge(10)
	p.Eat()
	fmt.Printf("%+v", p)
	//继承
	stuextend := Studentextend{}
	stuextend.New("shipeng", 2, "hebei")
	fmt.Println(stuextend)
	stuextend.Walk() //覆盖了personextend的方法
	stuextend.Eat()  //继承的方法
	/***************************************************************************************
	interface code*接口实现
	*/
	fmt.Println("***********************interface code*接口实现****************************")
	name := "A8L"
	aodi := AoDi{}
	aodi.Drive(name)

	mycar := MyCar{&aodi} //因为mycar结构体中是一个接口，所以初始化时，需要传入一个实现接口的结构体方法
	mycar.Drive("mycar")

	name = "X6"
	mycar = MyCar{&BMW{}}
	mycar.Drive(name)
	mycar.IDrive.Drive(name)
	fmt.Println()

	/*************************************************************************************
	interface code接口学习2：多态
	*/
	animal := Factory("duck")
	animal.Singgua()
	animal = Factory("goose")

	animal.Singgua()
	fmt.Printf("animal is %s", animal.Type())
	fmt.Println("")

	/***********************************************************************************
	pointercode指针学习
	*/
	//Point1()
	//Point2()
	//Point3()
	//Point4()
	Point5()

	/***********************************************************************************
	structcode指针学习
	*/
	Strunctfunc1()
	Structfunc2()
	Structfunc3()

	/***********************************************************************************
	compcode复合类型学习
	*/
	/*type关键字*/
	Compfunc()
	Compfunc2()
	/*struct*/
	Compstd2_func()
	Compstd2_func2()
	/*struct嵌套*/
	Compstd3_func()
	/*匿名struct 一般用来全局的数据库配置*/
	config.uid = "shipeng"
	config.pwd = "pwd"
	Compstd4_func()

	/*数组的使用*/
	//Compstd5_func()
	var arr [10]int
	length := len(arr)
	for i := 1; i <= length; i++ {
		arr[i-1] = i
	}
	Compstd5_func2(arr)  //用此函数修改arr元素的值
	fmt.Println(arr)     //仍然为[1 2 3 4 5 6 7 8 9 10]
	Compstd5_func3(&arr) //使用指针改变
	fmt.Println(arr)     //使用指针可以改变其值[2 4 6 8 10 12 14 16 18 20]

	/*compstd6切片的使用*/
	Compstd6_func()
	Compstd6_func2()
	Compstd6_func3()
	Compstd6_func4()

	/*compstd7字典使用*/
	Compstd7_func()
	Compstd7_func2()

	/*range关键字*/
	Compstd8_func()
	Compstd8_func2()
	/*接口函数的几种实现*/
	fmt.Println(Interfacefunstd())
	fmt.Println(Interfacefunstd2())
	GetFromSource(new(DB), "hello")
}
