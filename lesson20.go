// インターフェイス - 同じ型にして配列に入れることができる
// メソッド
// 構造体

package main

import "fmt"

type greeter interface {
	great()
}

type japanese struct{}

type american struct{}

func (j japanese) great() {
	fmt.Println("Konnichiwa!")
}

func (a american) great() {
	fmt.Println("hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.great()
	}
}
