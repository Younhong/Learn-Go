package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func main() {
	// name := "younhong"
	// fmt.Println(name)
	// fmt.Println(multiply(5, 7))
	// repeatMe(name, "yejin", "haeun", "sungjun")

	// totalLength, uppercase := lenAndUpper(name)
	// fmt.Println("Length:", totalLength, uppercase)

	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	// fmt.Println(canIDrink(16))

	// a := 2
	// b := &a
	// *b = 20
	// fmt.Println(a, *b)

	names := [5]string{"younhong", "april", "june"}
	names[3] = "minsu"
	names[4] = "yewon"
	fmt.Println(names)

	names2 := []string{"younhong", "april", "june"}
	names2 = append(names2, "yewon")
	fmt.Println(names2)
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 20 {
	// 	return false
	// }
	// return true
	switch {
	case age < 20:
		return false
	case age >= 20:
		return true
	}
	return false
}

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
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
