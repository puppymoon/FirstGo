package main

import "fmt"

func main() {
	// 创建一个程序，要求能够打印类似下面的结果（尾行达 25 个字符为止）：
	// G
	// GG
	// GGG
	// GGGG
	// GGGGG
	// GGGGGG
	c := "G"
	for i := 0; i < 26; i++ {
		fmt.Printf("%v\n", c)
		c += "G"
	}

}
