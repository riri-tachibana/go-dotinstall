// 配列 : スライス

package main

import "fmt"

func main() {
	// s := make([]int, 3) // [0 0 0]
	s := []int{1, 3, 5}
	// append - 配列に要素を追加
	s = append(s, 8, 2, 10)
	// copy
	t := make([]int, len(s))
	n := copy(t, s)
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(n)
}
