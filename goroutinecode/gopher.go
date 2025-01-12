package goroutinecode

import (
	"fmt"
	"sync"
	"time"
)

//wg.Done()  //用于表示当前携程执行结束
//职工
type Gopher struct {
	Name       string
	ID         int
	CoffeeName string
}

//两台咖啡机
var coffeeMachineArr = [2]CoffeeMachine{}

func init() {
	var coffeeMachine = CoffeeMachine{Name: "CoffeeMachine1"}
	var coffeeMachine2 = CoffeeMachine{Name: "CoffeeMachine2"}
	coffeeMachineArr[0] = coffeeMachine
	coffeeMachineArr[1] = coffeeMachine2
}

//make coff制作咖啡
func (this *Gopher) MakeCoffee(coffeeName string, wg *sync.WaitGroup) {
	//独有的咖啡机
	coffeeMachine := coffeeMachineArr[this.ID%2]
	//coffeeMachine枷锁
	coffeeMachine.Mlock.Lock()
	//(1）研磨咖啡（没有才研磨）
	if coffeeMachine.CoffeeName == "" {
		coffeeMachine.CoffeeName = coffeeName
		coffeeMachine.Gopher = *this
		fmt.Println("Gopher", this.ID, "Make Coffee", coffeeMachine.CoffeeName)
		time.Sleep(10 * time.Second)
	}
	//(2)倒咖啡
	this.TakeCoffee(coffeeMachine)
	//(3)喝咖啡
	this.DrinkCoffee(coffeeMachine)
	//释放锁
	coffeeMachine.Mlock.Unlock()
	wg.Done() //用于表示当前携程执行结束
}
func (this *Gopher) TakeCoffee(coffeeMachine CoffeeMachine) {
	if coffeeMachine.Name != "" {
		fmt.Println("Goper", this.ID, "Take coffee", coffeeMachine.Name)
		this.CoffeeName = coffeeMachine.CoffeeName
		time.Sleep(5 * time.Second)
		//倒完咖啡
		coffeeMachine.CoffeeName = ""
	} else {
		fmt.Println("Goper", this.ID, "Have no Coffee to  Take")
		this.CoffeeName = ""
	}
}

func (this *Gopher) DrinkCoffee(coffeeMachine CoffeeMachine) {
	if this.CoffeeName != "" {
		fmt.Println("Goper", this.ID, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Goper", this.ID, "Have no Coffee to Drink")
	}
}
