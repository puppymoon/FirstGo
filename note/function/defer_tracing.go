package main

import "fmt"

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	trace("a")
	// defer 為stack結構 先進後出
	// 所以會先印出leaving: a
	// 接著才是leaving: b
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
