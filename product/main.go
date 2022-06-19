package main

import (
	"fmt"
)

func main() {

	var a int
	var b int

	even := "Even"
	odd := "Odd"

	fmt.Scanf("%d %d", &a, &b)

	if a*b%2 == 0 {
		fmt.Println(even)
	} else {
		fmt.Println(odd)
	}
}
