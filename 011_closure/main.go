package main

import "fmt"

func main() {

	// result := func(a, b int) int {
	// 	return a + b
	// }(10, 20)

	// fmt.Println(result)
	closure("new")
}

//замыкание

func closure(s string) {
	text := "sentence " + "- " + s

	fmt.Println(text)

	result := func() string {
		fmt.Println(text + "some another text")
		return text + " new text"
	}

	fmt.Println(result())

}

// func sum(a, b int) int {

// 	func(s string) {
// 		fmt.Println(s)
// 	}("song")

// 	return a + b
// }
