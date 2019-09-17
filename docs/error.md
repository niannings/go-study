# 错误处理 (待完成)

```go
package main

import (
	"errors"
	"fmt"
)

// Error 是一个错误类型接口
type Error interface {
	Error() string
}

// Sqrt 是一个求平方根的函数
// 实现了Error接口
func Sqrt(f float64) (float64, Error) {
	if f < 0 {
		return 0, errors.New("请传入一个正数")
	}

	return 0, nil
}

func main() {
	_, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}
}
```