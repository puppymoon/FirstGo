package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

// 可以與fibonacci.go比較
func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacciCache(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fibonacciCache(n int) (res uint64) {
	// 若想要的n的值已經存在則直接返回
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciCache(n-1) + fibonacciCache(n-2)
	}
	// 將計算過的值存進對應n中
	fibs[n] = res
	return
}
