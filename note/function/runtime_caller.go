package main

import (
	"log"
	"runtime"
)

func f1() {
	where()
	log.Printf("now is f1")
}

func f2() {
	where()
	log.Printf("now is f2")
}

func f3() {
	where()
	log.Printf("now is f3")
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}
func main() {
	// 打開印出目前是哪一隻go
	log.SetFlags(log.Llongfile)
	log.Print("========>")
	var where2 = log.Print
	where()
	where2()
	f1()
	where()
	where2()
	f2()
	where()
	where2()
	f3()
}
