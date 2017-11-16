// 変数
// 変数名: １文字目が小文字＝そのpackageからだけ見える, 大文字だったら他のからもみれる

package main

import "fmt"

func main() {
	// var msg string // varで定義。変数名-型
	// msg = "hello world" // 定義したものを使える
	// var msg = "hello world" // 型は省略できる
	msg := "hello world"
	fmt.Println(msg)

	// var a, b int
	// a, b = 10, 15
	a, b := 10, 15

	var (
		c = int
		d = string
	)
	c = 20
	d = "hoge"
}
