package main

import "fmt"

// 假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？
func main() {
	str := "Google"
	str2 := Split2(str)
	fmt.Printf("The string %s transformed is: %s\n", str, str2)
}

func Split2(s string) string {
	mid := len(s) / 2
	return s[mid:] + s[:mid]
}
