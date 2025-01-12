package structcode

import (
	"fmt"
)

/************************************************************************
1 结构体的实例化和初始化
2 结构体接口实现
3 内嵌和组合
*/

type Person struct {
	Name  string
	Birth string
	ID    int64
}

func Strunctfunc1() {
	/*1结构体的实例化和初始化**************************************************/
	//1.1 第一种方法实例化
	var p1 Person //定义变量
	p1.Name = "王二小"
	p1.Birth = "1950.1.1"
	fmt.Println(p1)
	//1.2第二种方法实例化，函数实例化
	p2 := new(Person)
	p2.Name = "王二小"
	p2.Birth = "1950.1.1"
	fmt.Println(p2)
	//1.3 第三种方法实例化，取地址实例化
	p3 := &Person{}
	p3.Name = "王二小"
	p3.Birth = "1950.1.1"
	fmt.Println(p3)

	//1.4 结构体初始化
	p4 := Person{
		Name:  "wang",
		Birth: "1900-1-1",
	}
	fmt.Println(p4)
	//1.5 直接初始化
	p5 := &Person{
		"wangwu",
		"1900-1-1",
		5,
	}
	fmt.Println(p5)
}

/*******************2结构体接口实现*********************************************************/
//2.1 定义接口
type Cat interface {
	CatchMouse()
}
type Dog interface {
	Bark()
}

//2.2定义猫和狗公用结构体
type CatDog struct {
	Name string
}

//2.3 实现接口
func (catDog *CatDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n", catDog.Name)
}
func (catDog *CatDog) Bark() {
	fmt.Printf("%v barked loudly!\n", catDog.Name)
}

//2.4 统一函数函数调用
func Structfunc2() {
	//初始化一个结构体，并将其地址赋值给一个接口cat(重要)，即传入了func (catDog *CatDog)
	catDog := &CatDog{"lucky"}
	var cat Cat
	cat = catDog
	cat.CatchMouse()

	//声明一个dog接口，并将catDog指针赋值给dog
	var dog Dog
	dog = catDog
	dog.Bark()
}

/*******************3内嵌和组合*********************************************************/
//3.1游泳特性
type Swimming struct {
}

func (swim *Swimming) swim() {
	fmt.Println("swimming is my ability")
}

//3.2飞行特性
type Flying struct {
}

func (fly *Flying) fly() {
	fmt.Println("flying is my ability")
}

//3.3 野鸭子，具备飞行和游泳特性，内嵌上述两个结构体
type WildDuck struct {
	Swimming
	Flying
}

type DomestrcDuck struct {
	Swimming
}

//3.4定义函数统一调用
func Structfunc3() {
	//wildduck融合了两个结构体
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	domestrcDuck := DomestrcDuck{}
	domestrcDuck.swim()

	//测试
	strst := &WildDuck{}
	strst.fly()
	strst.swim()
}
