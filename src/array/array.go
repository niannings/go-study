package main

import "fmt"

/*
	## 声明数组
	一维数组
	var variable_name [SIZE] variable_type
	多维数组
	var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	## 初始化
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	自动大小
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 其实这里定义里一个切片
*/

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 二维数组 */
	// a := [3][4]int{
	// 	{0, 1, 2, 3},
	// 	{4, 5, 6, 7},
	// 	{8, 9, 10, 11}, /* 注意：这种写法末尾的逗号不可省 */
	// }

	println("指定大小：", add5array([5]int{1, 2, 3, 4, 5}))
	println("不定大小：", addArray([]int{1, 2, 3, 4, 5, 6}))
}

/* 向函数传递数组 */
// 1. 定义大小
func add5array(arr [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}

	return sum
}

// 2. 不定大小
func addArray(arr []int) int {
	size := len(arr)
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	return sum
}
