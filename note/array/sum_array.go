package main

import "fmt"

// a) 写一个 Sum() 函数，传入参数为一个 float32 数组成的数组 arrF，返回该数组的所有数字和。

// 如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？

// b) 写一个 SumAndAverage() 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。

func main() {
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	x := Sum(a)
	fmt.Printf("The sum of the array is: %f\n", x)
	sum, average := SumAndAverage(a)
	fmt.Printf("The sum of the array is: %f, and the average is: %f\n", sum, average)
	var b = []int{1, 2, 3, 4}
	sum1, average1 := SumAndAverageInt(b)
	// int / int 無條件捨去
	fmt.Printf("The sum of the array is: %d, and the average is: %d\n", sum1, average1)
}

func Sum(arrF []float32) (sum float32) {
	for _, value := range arrF {
		sum += value
	}
	return
}

func SumAndAverageInt(arrF []int) (sum, aver int) {
	for _, value := range arrF {
		sum += value
	}
	aver = sum / len(arrF)
	return
}

func SumAndAverage(arrF []float32) (sum, aver float32) {
	for _, value := range arrF {
		sum += value
	}
	aver = sum / float32(len(arrF))
	return
}
