package main

import "fmt"

func main() {
	// Add2返回一個func將其賦予給p2
	p2 := Add2()
	// 印出5
	fmt.Printf("Call Add 2 for 3 gives: %v\n", p2(3))
	// 一樣返回func 但其中a值已經被設定成2
	TwoAdder := Adder(2)
	// 2+3 印出5
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return 2 + b
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
