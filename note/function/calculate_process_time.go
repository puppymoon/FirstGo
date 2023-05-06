package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println(end)
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
