package main

import "fmt"

// 写一个函数 InsertStringSlice() 将切片插入到另一个切片的指定位置。
func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0) // at the front
	fmt.Println(res)                   // [A B C M N O P Q R]
	res = InsertStringSlice(s, in, 3)  // [M N O A B C P Q R]
	fmt.Println(res)
}

func InsertStringSlice(slice, insertion []string, index int) (res []string) {
	res = make([]string, len(slice)+len(insertion))
	copy(res, slice[:index])
	copy(res[index:], insertion)
	copy(res[index+len(insertion):], slice[index:])
	return
}
