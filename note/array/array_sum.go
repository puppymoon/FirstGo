package main

import "fmt"

// 一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种情况：
// 1. 传递数组的指针
// 2. 使用数组的切片
// 此為传递数组的指针

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	// 傳遞指針給Sum
	x := Sum(&array)
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	// 回傳 加總後的總和
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
