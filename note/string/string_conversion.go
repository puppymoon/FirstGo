package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	//目前環境int Size
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// 將字串轉換成int
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	// 將int轉換成字串所代表的數字（十進位）
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
