package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	// for ix, val := range coll { }
	// pos為索引 char為值
	for pos, char := range str {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %-5d  %U   '%c'   %X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
