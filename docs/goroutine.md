# 协程 goroutine (待完成)

```go
package main

/*
	## 同步
	### 基本概念
	1. 程序的初始化在一个独立的goroutine中执行。在初始化过程中创建的goroutine将在 第一个用于初始化goroutine执行完成后启动。
	2. 如果包p导入了包q，包q的init 初始化函数将在包p的初始化之前执行。
	3. 程序的入口函数 main.main 则是在所有的 init 函数执行完成 之后启动。
	4. 在任意init函数中新创建的goroutines，将在所有的init 函数完成后执行。
*/

var str string

func foo() {
	println(str)
}

func bar() {
	str = "hello go"
	go foo()
	// 会在某个时刻打印“hello go”（有可能是在bar函数返回之后）
}

func main() {
	bar()
}
```