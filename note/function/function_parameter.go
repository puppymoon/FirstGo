package main

import (
	"fmt"
)

func main() {
	// 將func當作參數
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(x int, f func(int, int)) {
	f(x, 2) // this becomes Add(1, 2)
}
