package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	// Fields 用空白切割字串
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("[%s], ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	// Split 使用指定字串切割
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("[%s], ", val)
	}
	fmt.Println()
	// 用指定字串將slice拼接起來
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}
