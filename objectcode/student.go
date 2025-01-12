package objectcode

import "fmt"

/*通过结构体组合person来实现继承person的属性和方法*/
type Studentextend struct {
	Personextend //组合personextend
	school       string
}

func (this *Studentextend) GotoSchool() {
	fmt.Println(this.name, " go to", this.school)
}

//覆盖了person.walk方法
func (this *Studentextend) Walk() {
	fmt.Println(this.name, " Walk")
}

func (this *Studentextend) New(name string, age int, school string) {
	//继承的属性
	this.age = age
	this.name = name
	this.school = school
}
