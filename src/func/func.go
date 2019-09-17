package main

func main() {
	println(add(1, 2))
	// 字符型可以和数字运算
	println(add('1', 'a'))
	/*
		函数的用法
		1. 高阶函数（就是一个函数作为另一个函数的参数）
		2. 闭包（函数内部的函数使用了外部函数的变量）
		3. 作为结构体类型的方法
		type Circle struct {
			radius float64
		}

		//该 method 属于 Circle 类型对象中的方法
		func (c Circle) getArea() float64 {
			//c.radius 即为 Circle 类型对象中的属性
			return 3.14 * c.radius * c.radius
		}
	*/
}

func add(a, b int) int {
	return a + b
}
