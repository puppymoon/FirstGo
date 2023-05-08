package main

import "fmt"

// 写一个函数 RemoveStringSlice() 将从 start 到 end 索引的元素从切片中移除。
func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	res := RemoveStringSlice(s, 2, 4)
	fmt.Println(res) // [M N Q R]
}

func RemoveStringSlice(slice []string, start, end int) (res []string) {
	res = make([]string, len(slice)-(end-start))
	copy(res, slice[:start])
	copy(res[start:], slice[end:])
	return
}
