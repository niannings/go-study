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

## 索引
- [x]  [变量](./docs/variable.md)
- [ ]  常量
