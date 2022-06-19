package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var arr []string

	fmt.Scanf("%s", &input)

	arr = strings.Split(input, "")

	num := 0
	for i := 0; i < 3; i++ {
		if arr[i] == "1" {
			num++
		}
	}
	fmt.Println(num)
}
