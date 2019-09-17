package main

/*
	## 理解 5个关键点
	1. interface 是一种类型
	2. interface 变量存储的是实现者的值
	3. 如何判断 interface 变量存储的是哪种类型
	4. 空的 interface
*/

import (
	"fmt"
)

// Phone 接口
type Phone interface {
	call()
}

// NokiaPhone 类型
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// IPhone 类型
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 4. 空的 interface
// doSomething将可以接受任何类型的参数
func doSomething(v interface{}) {
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
	/*
		不难看出 interface 的变量中存储的是实现了 interface 的类型的对象值，这种能力是 duck typing。
		在使用 interface 时不需要显式在 struct 上声明要实现哪个 interface ，只需要实现对应 interface 中的方法即可，
		go 会自动进行 interface 的检查，并在运行时执行从其他类型到 interface 的自动转换，即使实现了多个 interface，
		go 也会在使用对应 interface 时实现自动转换，这就是 interface 的魔力所在。
	*/

	// 3. 如何判断 interface 变量存储的是哪种类型
	if t, ok := phone.(*IPhone); ok {
		fmt.Println("IPhone implements Phone", t)
	}

	// 区分多种类型
	switch t := phone.(type) {
	case *IPhone:
		fmt.Println("i store *IPhone", t)
	case *NokiaPhone:
		fmt.Println("i store *NokiaPhone", t)
	}
}
