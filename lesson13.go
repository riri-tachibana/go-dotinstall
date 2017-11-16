// map - 連想配列

package main

import "fmt"

func main() {
	// m := make(map[string]int) // keyがstr,valueがint
	// m["taguchi"] = 200
	// m["fkoji"] = 300
	m := map[string]int{"taguchi": 100, "fkoji": 400}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "taguchi")
	v, ok := m["fkoji"]
	fmt.Println(v)
	fmt.Println(ok)
}
