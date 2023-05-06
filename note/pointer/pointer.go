package main

import "fmt"

func main() {
	// %d: digit   (10進位的數字)
	// %c: char    (字元)
	// %s: string  (字串)
	// %v: value   (值)
	var i1 = 5
	// 印出i1的位置
	fmt.Printf("An integer: %v, its location in memory: %p\n", i1, &i1)
	// 宣告一個指針變量
	var intP *int
	// 指針未分配到變量前 值為nil
	fmt.Printf("The value at memory location %v \n", intP)
	// 將指針指向i1
	intP = &i1
	// intP為值為指針的位置 *intP則為指針所指向的內容
	fmt.Printf("The value at memory location %v is %v\n", intP, *intP)
}
