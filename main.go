package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func main() {
	name := "younhong"
	// fmt.Println(name)
	// fmt.Println(multiply(5, 7))
	// repeatMe(name, "yejin", "haeun", "sungjun")

	totalLength, uppercase := lenAndUpper(name)
	fmt.Println("Length:", totalLength, uppercase)
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
