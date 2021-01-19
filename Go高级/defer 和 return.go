package main

import (
	"fmt"
	"time"
)

//  !!! defer 在 return 之前 !!!
//  !!! return 最后执行
func main()  {
	fmt.Println("f1 result: ", f1())
	fmt.Println("f2 result: ", f2())
}

func f1() int {
	var i int
	defer func() {
		i++
		time.Sleep(3 * time.Second)
		fmt.Println("f11: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f12: ", i)
	}()

	i = 1000
	return i
}

func f2() (i int) {
	defer func() {
		i++
		time.Sleep(2 * time.Second)
		fmt.Println("f21: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f22: ", i)
	}()

	i = 1000
	return i
}

//原因就是return会将返回值先保存起来，对于无名返回值来说，
//保存在一个临时对象中，defer是看不到这个临时对象的；
//而对于有名返回值来说，就保存在已命名的变量中。

//作者：JackieZheng
//链接：https://juejin.im/post/6844903877033066509
//来源：掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。