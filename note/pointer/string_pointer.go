package main

import "fmt"

func main() {
	s := "good bye"
	// 將指針p 指向 s
	var p *string = &s
	// 將指針的值改成ciao
	*p = "ciao"
	fmt.Printf("Here is the s location: %p\n", &s) // prints address
	fmt.Printf("Here is the pointer p: %p\n", p)   // prints address
	fmt.Printf("Here is the string *p: %s\n", *p)  // prints string
	fmt.Printf("Here is the string s: %s\n", s)    // prints same string
}
