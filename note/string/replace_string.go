package main

import (
	"fmt"
	"strings"
)

func main() {
	// using the function
	fmt.Println(strings.Replace("gfg gfg gfg", "g", "G", 2))
	fmt.Println(strings.Replace("gfg gfg gfg", "fg", "FG", -1))
}
