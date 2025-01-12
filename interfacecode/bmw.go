package interfacecode

import "fmt"

type BMW struct{}

func (*BMW) Drive(name string) {
	fmt.Println("Drive BMW  " + name + " Car")
}
