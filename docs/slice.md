# 切片（动态数组）

## 基本概念
1. Go 语言切片是对数组的抽象。
2. Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

## 定义切片
1. 使用未指定大小的数组定义：```var identifier []type```
2. 使用```make```：```var slice1 []type = make([]type, len)```
3. ```make```第三个参数指定切片最大容量：```make([]T, length, capacity)```

## 切片初始化
```go
s :=[] int {1,2,3 }
```
具体看代码

## ```len()``` 和 ```cap()``` 函数
1. 切片是可索引的，并且可以由 len() 方法获取长度。
2. 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

## 空(```nil```)切片
一个切片在未初始化之前默认为 nil，长度为 0

## ```append()``` 和 ```copy()``` 函数
1. 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	arr := [3]int{4, 5, 6}
	// 初始化切片s1,是数组arr的引用
	s1 := arr[:]

	fmt.Printf("s = %v\n", s)
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("s1 = %v\n", s1)

	// 截取 1 ~ 末尾 （包括 1）
	s2 := arr[1:]
	// 截取 开头 ~ 1 （不包括 1）
	s3 := arr[:1]

	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s3 = %v\n", s3)

	var numbers []int

	println("\n查看 numbers的信息")
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("大小=%d 最大容量=%d 切片值=%v\n", len(x), cap(x), x)
}
```
