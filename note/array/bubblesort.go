package main

import "fmt"

// 编写一个程序，使用冒泡排序的方法排序一个包含整数的切片
func main() {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	bubbleSort(sla)
	fmt.Println("after sort:  ", sla)
}

func bubbleSort(sl []int) {
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				temp := sl[j]
				sl[j] = sl[j+1]
				sl[j+1] = temp
			}
		}
	}
}
