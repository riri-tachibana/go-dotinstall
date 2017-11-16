// 基本的なデータ型
/*
string "hello"
int     53
float64 10.2
bool     false true
nil

var s string // "" nilではない
var a int // 0 nilではない
var f bool // false nilではない
*/

package main

import "fmt"

func main() {
	a := 10
	b := 12.3
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d) // %dは整数値、%bは小数、%sは文字列、%tは真偽値
}
