package main

import (
	"fmt"
)

// 將ab相乘後的值賦予給reply指標的位置
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	// reply為n指標的位置
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
