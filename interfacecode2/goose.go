package interfacecode2

import "fmt"

type Goose struct {
	Color string
	Age   int
}

func (this *Goose) Sleep() {
	fmt.Println(this.Color + " goose sleep")
}
func (this *Goose) Singgua() {
	fmt.Println(this.Color + " goose Singgua")
}
func (this *Goose) Type() string {
	return "Goose"
}
