// インターフェイス - 同じ型にして配列に入れることができる
// メソッド
// 構造体

package main

import "fmt"

func show(t interface{}) {
	//型アサーション
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("i am jpn")
	// } else {
	// 	fmt.Println("i am not jpn")
	// }

	//型switch
	switch t.(type) {
	case japanese:
		fmt.Println("i am jpn")
	default:
		fmt.Println("i am not jpn")
	}
}

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
		show(greeter)
	}
}
