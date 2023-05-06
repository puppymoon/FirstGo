package main

import "fmt"

func main() {
	test := fib()
	for i := 0; i <= 10; i++ {
		result := test(i)
		fmt.Printf("fib(%d) is: %d\n", i, result)
	}
}

func fib() func(int) int {
	a, b := 1, 1
	return func(n int) (res int) {
		if n == 0 {
			res = a
		} else if n == 1 {
			res = b
		} else {
			a, b = b, a+b
			res = b
		}
		return
	}
}
