package main

import "fmt"

func main() {

	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	for ix, season := range seasons {
		// 此處的season為此for的copy過來局部變量
		// 對其進行變更 不會影響原slice的值
		if ix == 2 {
			season = "XXXXX"
		}
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	// 不需要index的用法
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	// 不需要value的用法
	for ix := range seasons {
		fmt.Printf("%d", ix)
	}
}
