package main

import "fmt"

/*
	## 指针声明
	- 一个指针变量指向了一个值的内存地址。
	- 类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
	```var var_name *var-type```
	例如：
	```go
		var ip *int        // 指向整型
		var fp *float32    // 指向浮点型
	```

	## 空指针
	1. 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	2. nil 指针也称为空指针。
	3. nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	## 指针数组（和C一样）
	## 多级指针（和C一样）
	## 作为参数的指针（和C一样）见末尾swap函数
*/

func main() {
	var a int = 10
	// 取地址
	var ip *int = &a

	// 取值
	fmt.Printf("变量的地址: %x\n变量的值是：%d\n", ip, *ip)
	var ptr, b *int

	// ptr 的值为 : 0
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	// 判断是否是空指针
	println(ptr == nil)
	// 空指针指向的地址是： 0x0 , 0x0
	println("空指针指向的地址是：", ptr, ",", b)
	// ptr 的值为 : <nil>
	fmt.Println("ptr 的值为 : ", ptr)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
