# 协程 goroutine

## 同步
### 基本概念
1. 程序的初始化在一个独立的```goroutine```中执行。在初始化过程中创建的```goroutine```将在 第一个用于初始化```goroutine```执行完成后启动。
2. 如果包p导入了包```q```，包```q```的```init``` 初始化函数将在包```p```的初始化之前执行。
3. 程序的入口函数 ```main.main``` 则是在所有的 ```init``` 函数执行完成 之后启动。
4. 在任意```init```函数中新创建的```goroutines```，将在所有的```init``` 函数完成后执行。

```go
package main

import "time"

var str string
var complete chan string = make(chan string)

func foo() {
	println(str)
	complete <- "我听到了。妹子～，咦，妹子在哪呢？" //
}

func bar() {
	str = "你看不见我"
	go foo()
	/*
		Q: 这里开启一个goroutine去执行foo，为何没有输出任何东西 ?
		A: 因为这个goroutine还没来得及跑，bar函数就已经退出了
	*/
}

func baz() {
	str = "我在这儿呢"
	go foo()
	time.Sleep(time.Second) // 停顿一秒
	/*
		把main函数里的 bar() 改成 baz() 就可以输出了
		但是采用等待的方式并不好，最好是子goroutine告诉主goroutine我要跑完了
		那么信道channel可以做到阻塞主goroutine
	*/
}

func niceBaz() {
	str = "有妹子喊你"
	go foo()
	// println(<-complete)
}

func main() {
	// bar()
	// baz()
	niceBaz()
	// complete <- "哈哈" // 数据流入无缓冲信道, 如果没有其他goroutine来拿走这个数据，那么当前线阻塞
}
```