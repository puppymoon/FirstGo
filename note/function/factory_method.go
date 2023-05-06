package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// 返回一個suffix固定為.bmp的func(name string) string
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file.bmp"))
	fmt.Println(addBmp("file.bm"))
	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("file"))
}
