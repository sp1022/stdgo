package objectcode

import "fmt"

/*面向对象************************************
通过结构体封装实现
*/
type Personobj struct {
	name string
	age  int
}

func (this *Personobj) SetName(name string) {
	this.name = name
}
func (this *Personobj) SetAge(age int) {
	this.age = age
}
func (this *Personobj) GetName() string {
	return this.name
}
func (this *Personobj) GetAge() int {
	return this.age
}

//方法大写才能被包外使用
func (p *Personobj) Eat() {
	fmt.Println("person can eat")
}
