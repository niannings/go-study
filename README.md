# Golang 学习笔记
## 安装Golang
1. 下载Go[下载地址](https://golang.org/dl/)
2. 初始化
```bash
mkdir goland & cd goland
touch test.go & code hello.go
```
3. 编写第一个go程序 hello world
```go
// hello.go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
4. 运行
```bash
go run hello.go
# hello world
```
5. 构建应用程序
```bash
go build
# 接下来运行
./goland
# hello world
```

## 目录索引
- [x]	[基础语法](https://www.runoob.com/go/go-data-types.html)
- [x]  	[数据类型](https://www.runoob.com/go/go-data-types.html)
- [x]  	[变量](./docs/variable.md)
- [x] 	[常量](./docs/constant.md)
- [x] 	[运算符](https://www.runoob.com/go/go-operators.html) 注意：自增/减不支持前置，如```++i```，```--i```
- [x]	条件语句（比C多了个select语句）；循环语句（比C少了while循环）
- [x]	[函数](./docs/func.md)
- [x]	[作用域](./docs/scope.md)
- [x]	[数组](./docs/array.md)
- [x]	[指针](./docs/pointer.md)
- [x]	[结构体](./docs/struct.md)
- [x]	[切片](./docs/slice.md)
- [x]	[范围(range)](./docs/range.md)
- [x]	[Map](./docs/Map.md)
- [x]	[接口](./docs/interface.md)
- [ ]	[错误处理](./docs/error.md)
- [x]	[协程](./docs/goroutine.md)
- [x]	[信道](./docs/channel.md)

## 参考资料
- [go官方文档](https://golang.org/doc/)
- [菜鸟go](https://www.runoob.com/go/go-tutorial.html)
- [go中文文档](http://shouce.jb51.net/golang-doc)
- [理解 Go interface 的 5 个关键点](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)