package objectcode

import "fmt"

/*面向对象，组合实现继承*/
type Personextend struct {
	name string
	age  int
}

func (p *Personextend) Walk() {
	fmt.Println("person walk")
}

func (p *Personextend) Eat() {
	fmt.Println("person Eat")
}
