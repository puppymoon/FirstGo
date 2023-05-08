package main

import "fmt"

func copyArr(a [3]int) {
	// 複製的ar位置
	fmt.Printf("%p\n", &a)
}

func main() {
	var ar [3]int
	// 一開始的ar位置
	fmt.Printf("%p\n", &ar)
	copyArr(ar)

}
