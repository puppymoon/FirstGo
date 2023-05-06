package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	// x 為func內部的變數 隨著每次執行增加
	var x int
	return func(delta int) int {
		// 第一次為0 第二次為1 第三次為21
		fmt.Printf("x=====>%v\n", x)
		x += delta
		return x
	}
}
