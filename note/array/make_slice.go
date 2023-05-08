package main

import "fmt"

func main() {
	// make 做出一個slice type:int length:10 (cap:10)
	var slice1 []int = make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 问题 7.3 给定 s := make([]byte, 5)，len(s) 和 cap(s) 分别是多少？s = s[2:4]，len(s) 和 cap(s) 又分别是多少？
	t := make([]byte, 5)
	fmt.Println("len(t):%v", len(t))
	fmt.Println("cap(t):%v", cap(t))
	t = t[2:4]
	fmt.Println("len(t):%v", len(t))
	fmt.Println("cap(t):%v", cap(t))
	// 问题 7.4 假设 s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := s1[2:]，s2 的值是多少？如果我们执行 s2[1] = 't'，s1 和 s2 现在的值又分别是多少？

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("s1:%c\n", s1) // ['p', 'o', 'e', 'm']
	fmt.Printf("s2:%c\n", s2) // ['e', 'm']
	s2[1] = 't'
	fmt.Printf("s1:%c\n", s1) // ['p', 'o', 'e', 't']
	fmt.Printf("s2:%c\n", s2) // ['e', 't']
}
