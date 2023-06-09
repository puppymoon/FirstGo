package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	slice1 = slice1[0:0]
	fmt.Println(len(slice1))
	slice1 = slice1[0:1]
	fmt.Println(len(slice1))
	// 不能超過原本的cap
	// panic: runtime error: slice bounds out of range [:11] with capacity 10
	slice1 = slice1[0:11]
}
