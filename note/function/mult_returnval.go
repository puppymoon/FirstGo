package main

import (
	"fmt"
)

// 使用非命名返回值是很糟的编程习惯
func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

// 建議寫法
func SumProductDiffN(i, j int) (s int, p int, d int) {
	s = i + j
	p = i * j
	d = i - j
	return
}

func main() {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum, prod, diff = SumProductDiffN(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
}
