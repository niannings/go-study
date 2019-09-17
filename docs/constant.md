# 常量

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ## 常量声明和变量类似
	const a int = 1
	fmt.Printf("常量a = %d\n", a)
	// ## 常量还可以用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	// ## 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
	const (
		b = "abc"
		c = len(b)
		d = unsafe.Sizeof(a)
	)
	// 输出 abc 3 16
	println(a, b, c)
	// ## iota
	/*
		1. iota，特殊常量，可以认为是一个可以被编译器修改的常量。
		2. iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
		const (
			a = iota // 0
			b = iota // 1
			c = iota // 2
		)
		3. 用法
		const (
			a = iota   //0
			b          //1
			c          //2
			d = "ha"   //独立值，iota += 1
			e          //"ha"   iota += 1
			f = 100    //iota +=1
			g          //100  iota +=1
			h = iota   //7,恢复计数
			i          //8
		)
	*/
	// ## 一个有趣的位运算例子
	const (
		i = 1 << iota // 1 << 0		1 -> 1 		= 1
		j = 3 << iota // 3 << 1		11 -> 110 	= 6
		k             // 3 << 2		11 -> 1100	= 12
		l             // 3 << 3		11 -> 11000	= 24
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
```