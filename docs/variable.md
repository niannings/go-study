```go
package main

import (
	"fmt"
)

func main() {
	// hello world
	fmt.Printf("hello world\n")
	// 数据类型
	// ## 一、变量声明
	// 变量声明后必须使用，否则编译报错
	var a int
	var b, c int
	// 输出0 0 0 数字的初始值为0
	fmt.Println(a, b, c)
	/*
		1. 数值类型（包括complex64/128）为 0
		2. 布尔类型为 false
		3. 字符串为 ""（空字符串）
		4. 以下几种类型为 nil：
			var a *int
			var a []int
			var a map[string] int
			var a chan int
			var a func(string) int
			var a error // error 是接口
	*/
	// ## 二、自动判断变量类型
	var d = true
	// 输出 true
	fmt.Println(d)
	// ## 三、省略var
	e := "girl"
	fmt.Println(e)
	/*
		注意 := 左侧如果没有声明新的变量，就产生编译错误
		var intVal int
		intVal :=1 // 这时候会产生编译错误
		intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，**因为 := 是一个声明语句**
		这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
	*/
	// ## 四、多变量声明
	/*
		//类型相同多个变量, 非全局变量
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3
		var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
		vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

		// 这种因式分解关键字的写法一般用于声明全局变量
		var (
			vname1 v_type1
			vname2 v_type2
		)
	*/

	/*
		其他笔记📒
		a, b, c := 5, 7, "abc"
		1. 右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。
		2. 这被称为 并行 或 同时 赋值。
		3. 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
		4. 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
		5. _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
		6. 并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
	*/
}
```