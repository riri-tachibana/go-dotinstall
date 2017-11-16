// 関数

package main

import "fmt"

// func hi(name string) string {
// 	// fmt.Println("hi!" + name)
// 	msg := "hi" + name
// 	return msg
// }

// func main() {
// 	// hi("taguchi")
// 	fmt.Println(hi("taguchi"))
// }

func hi(name string) (msg string) {
	msg = "hi!" + name
	return
}

func main() {
	fmt.Println(hi("taguchi"))
}
