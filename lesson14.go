// if
/*
> >=
< <=
==
!=
&& || !
*/

package main

import "fmt"

func main() {
	// score := 83

	if score := 43; score > 80 {
		fmt.Println("great")
	} else if score > 60 {
		fmt.Println("good")
	} else {
		fmt.Println("soso")
	}
}
