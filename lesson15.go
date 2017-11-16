// switch

package main

import "fmt"

func main() {
	// signal := "red"
	// switch signal {
	// case "red":
	// 	fmt.Println("stop")
	// case "yellow":
	// 	fmt.Println("caution")
	// case "green", "blue":
	// 	fmt.Println("go")
	// default:
	// 	fmt.Println("wrong signal")
	// }

	score := 82
	switch {
	case score > 80:
		fmt.Println("Great")
	default:
		fmt.Println("soso")
	}
}
