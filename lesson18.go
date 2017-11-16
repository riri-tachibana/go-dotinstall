// 構造体

package main

import "fmt"

type user struct {
	name  string //初期化では空文字
	score int    // 初期化では0
} //fieldという

func main() {
	// u := new(user) //初期化
	// // (*u).name = "taguchi"
	// u.name = "taguchi"
	// u.score = 20

	// u := user{"taguchi", 50}
	u := user{name: "taguchi", score: 50}
	fmt.Println(u)
}
