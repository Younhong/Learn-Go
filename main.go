package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	// go count("younhong")
	// count("yerin")
	people := [4]string{"younhong", "yerin", "minki", "yewon"}
	for _, person := range people {
		go isCute(person, c)
	}
	// result := <-c
	fmt.Println("Waiting for messages")
	for i := 0; i < len(people); i++ {
		fmt.Println("Received Message:", <-c)
	}
	// time.Sleep(time.Second * 10)
}

// func count(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is blah-blah", i)
// 		time.Sleep(time.Second)
// 	}
// }

func isCute(person string, c chan string) {
	// fmt.Println(person)
	time.Sleep(time.Second * 10)
	c <- person + " is cute"
}
