package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	// return前 將返回值印出
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}
