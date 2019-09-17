package main

/*
	作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

	Go 语言中变量可以在三个地方声明：
	1. 函数内定义的变量称为局部变量
	2. 函数外定义的变量称为全局变量
	3. 函数定义中的变量称为形式参数
*/

// 全局变量
var a int = 20

func main() {
	// 局部变量
	var a int = 30
	println(a)
}

func add(a, b int) int {
	// a, b为函数add的形式参数，作为该函数的局部变量
	return a + b
}
