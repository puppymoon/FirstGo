package main

import "fmt"

func main() {
	var arr1 = []int{0, 1, 2, 3, 4, 5}
	// 從2開始 [ 2, 3, 4] 不包含5
	var slice1 []int = arr1[2:5]

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// slice 有三個參數 ( pointer, length, capacity)
	// slice1的 length為3 capacity為4
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 與array不同 slice可以變動尺寸
	// [ 2, 3, 4, 5]
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	// 超出原本的範圍
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	var slice2 []byte = b[1:4]
	fmt.Printf("b[1:4]: %c\n", slice2)
	slice2 = b[:2]
	fmt.Printf("b[:2]: %c\n", slice2)
	slice2 = b[2:]
	fmt.Printf("b[2:]: %c\n", slice2)
	slice2 = b[:]
	fmt.Printf("b[:]: %c\n", slice2)
}
