package main

import "fmt"

// 编写一个程序，要求能够反转字符串，即将 "Google" 转换成 "elgooG"（提示：使用 []byte 类型的切片）。
// 如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。
func main() {
	str := "Google"
	fmt.Println(str, " --> ", reverse(str))
}

func reverse(s string) (res string) {
	runes := []rune(s)
	i := 0
	length := len(runes)
	for i < len(runes) {
		temp := runes[i]
		runes[i] = runes[length-1]
		runes[length-1] = temp
		i++
		length--
		if i >= length {
			break
		}
	}
	res = string(runes)
	return
}
