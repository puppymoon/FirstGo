package main

import "fmt"

// 写一个 minSlice() 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice() 方法返回最大值。
func main() {
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	fmt.Printf("The maximum is %d\n", max)
	min := minSlice(sl1)
	fmt.Printf("The minimum is %d\n", min)
}

func minSlice(arr []int) (min int) {
	min = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return
}
func maxSlice(arr []int) (max int) {
	max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return
}
