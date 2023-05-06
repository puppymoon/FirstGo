package main

import "fmt"

func main() {

	//標籤有區分大小寫 一般來說都用全大寫
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			// 每當j = 4 時跳轉回上層for迴圈 故 j又重置為0
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
