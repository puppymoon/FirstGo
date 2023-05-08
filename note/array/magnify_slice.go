package main

import "fmt"

// 给定一个切片 s []int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。
func main() {
	var s []int
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	newSlice := make([]int, len(s)*factor)
	copy(newSlice, s)
	s = newSlice
	return s
}
