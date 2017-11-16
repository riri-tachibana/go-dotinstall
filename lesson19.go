// 構造体
// メソッド(データ型に紐づいた関数)

package main

import "fmt"

type user struct {
	name  string //初期化では空文字
	score int    // 初期化では0
} //fieldという

func (u user) show() {
	fmt.Printf("name:%s, score: %d\n", u.name, u.score)
} //レシーバー

func (u *user) hit() {
	u.score++
} // 参照渡し

func main() {
	u := user{name: "taguchi", score: 50}
	// fmt.Println(u)
	u.hit()
	u.show()
}
