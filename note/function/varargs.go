package main

import "fmt"

func F1(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	F1("test1", "test2", "test5566")
	F1("test1", "test2", "test5566", "test2", "test5566")

}
