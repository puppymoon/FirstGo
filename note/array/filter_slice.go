package main

import "fmt"

// 用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，
// 第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。
// 通过 fn 测试方法测试当整型值是偶数时的情况。
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
}

func Filter(s []int, fn func(int) bool) (res []int) {
	res = []int{}
	for _, value := range s {
		if fn(value) {
			res = append(res, value)
		}
	}
	return
}

func even(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
