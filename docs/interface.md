# 接口 interface (待完成)

## 理解 5个关键点
1. interface 是一种类型
2. interface 变量存储的是实现者的值
3. 如何判断 interface 变量存储的是哪种类型
4. 空的 interface
5.interface 的实现者的 receiver 如何选择

```go
package main

/*
	## 理解 5个关键点
	1. interface 是一种类型
	2. interface 变量存储的是实现者的值
	3. 如何判断 interface 变量存储的是哪种类型
	4. 空的 interface
	5.interface 的实现者的 receiver 如何选择
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

/*
	go 不会对 类型是interface{} 的 slice 进行转换
	参考官方wiiki https://github.com/golang/go/wiki/InterfaceSlice 大意是 interface{} 会占用两个字长的存储空间，
	一个是自身的 methods 数据，一个是指向其存储值的指针，也就是 interface 变量存储的值，
	因而 slice []interface{} 其长度是固定的N*2，但是 []T 的长度是N*sizeof(T)，两种 slice 实际存储值的大小是有区别的
*/
func printSlice(slice []interface{}) {
	fmt.Printf("%v", slice)
}

// Animal interface
type Animal interface {
	Speak() string
}

// Dog struct
type Dog struct {
}

// Speak is a method
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat struct
type Cat struct {
}

// Speak is a method
func (c *Cat) Speak() string {
	return "Meow!"
}

// Llama struct
type Llama struct {
}

// Speak is a method
func (l Llama) Speak() string {
	return "?????"
}

// JavaProgrammer struct
type JavaProgrammer struct {
}

// Speak is a method
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
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

	dataSlice := []int{1, 2, 3, 4, 5}
	// printSlice(dataSlice) // 会报错
	// 手动转换
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	printSlice(interfaceSlice) // 转化后即可传入
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}

	/*	5.interface 的实现者的 receiver 如何选择
		Q：	下面注释的代码取消注释后会报错：
			cannot use Cat literal (type Cat) as type Animal in array or slice literal:
			Cat does not implement Animal (Speak method has pointer receiver)
		A：	Animal的实现者Cat是个指针类型，与其他（值类型）不一致；值是拷贝，指针是引用。
		Q：	如果实现者是值类型，那么该方法在使用时可传值也可传指针，为什么？
		A：	如果是按 pointer 调用，go 会自动进行转换，因为有了指针总是能得到指针指向的值是什么，如果是 value 调用，go 将无从得知 value 的原始值是什么，
			因为 value 是份拷贝。go 会把指针进行隐式转换得到 value，但反过来则不行。

	*/
	// animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }
}
```
