package compcode

import "fmt"

/*3 struct嵌套********************************************
 */
type Address struct {
	City string
}

type student struct {
	name    string
	xh      string
	address Address
}

func Compstd3_func() {
	stu := student{}
	stu.name = "xiaos"
	stu.xh = "000"
	stu.address.City = "hangzhou"
	fmt.Printf("%+v", stu)
}
