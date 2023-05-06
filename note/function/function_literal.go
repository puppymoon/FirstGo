package main

import "fmt"

func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		// 匿名func直接賦值給變數
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
