package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	var s string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)

	fmt.Println(a+b+c, s)
}
