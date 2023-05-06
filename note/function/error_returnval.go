package main

import (
	"errors"
	"fmt"
	"math"
)

func MySqrt(i float64) (s float64, e error) {
	if i < 0 {
		return float64(math.NaN()), errors.New("value is below zero")
	}
	return math.Sqrt(i), nil
}

func MySqrtN(i float64) (s float64, e error) {
	if i < 0 {
		s = math.NaN()
		e = errors.New("value is below zero")
		return
	}
	s = math.Sqrt(i)
	e = nil
	return
}
func main() {
	ret1, err1 := MySqrt(5)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Println(MySqrtN(5))
}
