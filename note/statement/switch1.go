package main

import "fmt"

func main() {
	var num1 int = 100

	//與java不同 不用使用break 當num=98時 並不會進入case 100
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}
