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
- [ ]	切片