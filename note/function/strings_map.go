package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}

func asciiOnly(c rune) (x rune) {
	x = c
	if c > 127 {
		fmt.Printf("c > 127 char: %c\n", c)
		x = '?'
	}
	return
}
