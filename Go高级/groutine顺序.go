package main

import (
	"fmt"
	"time"
)

func main() {
	for i:= 0; i < 3; i++ {

		// 上面的示例中：
		// foo goroutine被 a 阻塞，
		// bar goroutine被 b 阻塞，
		// dog goroutine被 c 阻塞。
		// dog 依赖的c由 bar 关闭，bar 依赖的 b 由 foo 关闭。
		// 如此一来，当main goroutine中的 a 被关闭后，
		// A()从阻塞中释放，继续执行，关闭 b，然后B从阻塞中释放，继续执行，关闭z，C得以释放。
		// 由于z被关闭后，z仍然可读，所以多次执行C(z)不会出问题。

		// A()和B()不能多次执行，因为close()不能操作已被关闭的channel。

		// 注意，上面的channel都是struct{}类型的，整个过程中，x、y、z这3个通道都没有传递数据，
		// 而是直接关闭来释放通道，
		// 让某些阻塞的goroutine继续执行下去。显然，这里的x、y、z的作用都是"信号通道"，用来传递消息。

		a := make(chan struct{})
		b := make(chan struct{})
		c := make(chan struct{})

		// 控制循环
		d := make(chan struct{})
		go foo(a, b)
		go bar(b, c)
		go dog(c, d)
		
		// 控制启动
		close(a)

		// 控制循环
		<- d
	}
	time.Sleep(10 * time.Second)
}

func foo(a chan struct{}, b chan struct{}) {
	<- a
	fmt.Println("foo")
	close(b)
}

func bar(b chan struct{}, c chan struct{}) {
	<- b
	fmt.Println("bar")
	close(c)
}

func dog(c chan struct{}, d chan struct{}) {
	<- c
	fmt.Println("dog")
	close(d)
}