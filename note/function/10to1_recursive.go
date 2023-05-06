package main

import "fmt"

func main() {
	printrec(10)
}

func printrec(i int) {
	fmt.Println(i)
	i--
	if i == 0 {
		return
	}
	printrec(i)
}
