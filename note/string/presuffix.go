package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("Does the string \"%s\"\n", str)
	//Th開頭
	fmt.Printf("have prefix %s? %t\n", "Th", strings.HasPrefix(str, "Th"))
	//th結尾
	fmt.Printf("have prefix %s? %t\n", "Th", strings.HasSuffix(str, "Th"))
	//包含an
	fmt.Printf("contains %s? %t\n", "an", strings.Contains(str, "an"))
	//第一次出現的index
	fmt.Printf("'is' index %s %v\n", "is", strings.Index(str, "is"))

}
