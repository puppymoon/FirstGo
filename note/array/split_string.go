package main

import "fmt"

// 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func main() {
	str := "Google"
	for i := 0; i <= len(str); i++ {
		a, b := Split(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}
}

func Split(str string, i int) (str1 string, str2 string) {
	str1 = str[:i]
	str2 = str[i:]
	return
}
