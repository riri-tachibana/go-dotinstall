// ポインタ
// アドレスを指し示す変数
// 演算はできない

package main

import "fmt"

func main() {
	a := 5
	var pa *int //int型を格納するpaという領域ができたよ
	pa = &a     // &a = aのアドレス
	// paの領域にあるデータの値 = *pa
	fmt.Println(pa)
	fmt.Println(*pa)
}
