package interfacecode2

import "fmt"

type Duck struct {
	Color string
	Age   int
}

func (this *Duck) Sleep() {
	fmt.Println(this.Color + " Duck sleep")
}

func (this *Duck) Singgua() {
	fmt.Println(this.Color + " Duck Singgua")
}

func (this *Duck) Type() string {
	return "Duck"
}
