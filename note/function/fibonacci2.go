package main

import "fmt"

func main() {
	// result := 0
	for i := 0; i <= 10; i++ {
		position, result := fibonacci2(i)
		fmt.Printf("fibonacci2(%d) is: %d\n", position, result)
	}
}

func fibonacci2(n int) (pos, res int) {
	pos = n
	if n <= 1 {
		res = 1
	} else {
		_, res1 := fibonacci2(n - 1)
		_, res2 := fibonacci2(n - 2)
		res = res1 + res2
	}
	return
}
