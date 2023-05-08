package main

import "fmt"

func main() {
	var ar [16]int
	for i := 0; i < len(ar); i++ {
		ar[i] = i
	}
	fmt.Print(ar)
}
