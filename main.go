package main

import (
	"fmt"

	"github.com/younhong/Learn-Go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary["hello"])
}
