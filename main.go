package main

import (
	"fmt"

	"github.com/younhong/Learn-Go/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{"first": "First Word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	dictionary := mydict.Dictionary{}
	// err := dictionary.Add("hello", "Greeting")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// definition, err := dictionary.Search("hello")
	// fmt.Println(definition)
	// err2 := dictionary.Add("hello", "Greeting")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	dictionary.Add("hello", "first")
	// err := dictionary.Update("hello", "second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search("hello")
	dictionary.Delete("hello")
	word2, error := dictionary.Search("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(word2)
	}
}
