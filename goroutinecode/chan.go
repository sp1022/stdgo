/***********************************
chan 在 Golang 中是一种非常有用的并发机制，可以用于在 goroutine 之间传递数据。以下是一些 chan 的具体用处：

1 实现生产者-消费者模型
在生产者-消费者模型中，一个或多个生产者生成数据，一个或多个消费者消费这些数据。
chan 可以非常方便地实现这个模型，生产者将数据发送到 chan 中，而消费者从 chan 中接收数据。
使用 chan 可以避免显式的锁和条件变量，简化了代码的编写和理解。

2 控制并发
在 Golang 中，每个 goroutine 都是轻量级的线程，可以同时运行数千个 goroutine。
但是，同时运行太多的 goroutine 可能会导致系统的负载过高，从而影响系统的性能。
chan 可以用来控制并发，通过控制 chan 中的缓冲区大小，可以限制同时运行的 goroutine 数量，从而避免系统的负载过高。

3 实现任务调度
chan 可以用来实现任务调度。例如，在一个任务队列中，任务的执行顺序可能是随机的。
使用 chan 可以将任务放入一个 chan 中，然后使用 go 关键字启动多个 goroutine 来执行任务，
每个 goroutine 从 chan 中接收任务并执行。这样可以保证任务按照放入 chan 的顺序执行，从而实现任务的调度。

4 需要注意以下几个点：
(1)chan 中的数据类型必须是可比较的，也就是说，不能使用函数、映射或切片等类型作为 chan 的元素类型。
(2)当使用 chan 传递指针时，需要注意内存的安全。如果在一个 goroutine 中修改了指针所指向的内存，
而另一个 goroutine 仍在使用该指针，就可能会导致数据不一致的问题。
(3)使用 chan 时要注意避免死锁的情况。当一个 goroutine 在等待另一个 goroutine 的消息时，
如果两个 goroutine 都在等待对方发送消息，就会导致死锁的情况。需要设计好 chan 的使用方式，避免出现死锁的情况。
*************************************/
/*
信道分为三种：
(1) 只读信道：ch:=make(<- 类型)
(2)只写信道：ch:=make(chan <- 类型)
(3)双向信道:ch:=make(chan 类型)
*/
package goroutinecode

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Printf("Producer produced: %d\n", num)
		ch <- num
		time.Sleep(time.Second)
	}
	close(ch)
}

func Consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consumer consumed: %d\n", num)
		time.Sleep(time.Second)
	}
}
