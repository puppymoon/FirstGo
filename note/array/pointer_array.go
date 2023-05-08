package main

import "fmt"

func f(a [3]int) {
	a[1] = 2
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[2] = 3
	fmt.Println(a)
}

func main() {
	var ar [3]int
	// 傳遞一個複製過後的ar給f() 故在f()內變更值 不會影響原先的ar
	f(ar)
	// 傳遞指針給fp 更改後會影響原先ar的值
	fp(&ar) // passes a pointer to ar
	fmt.Println(ar)
}
