package main

import (
	"fmt"
)

func main() {

	// var array [5]int

	// array[0] = 111
	// array[1] = 222
	// array[2] = 333
	//fmt.Println(array)

	// for i := 0; i < len(array); i++ {
	// 	fmt.Println(i)
	// }

	// for i, v := range array {
	// 	fmt.Printf("Index - %v, Value - %v\n", i, v)
	// }

	var array2 = [...]string{"a", "b"}
	fmt.Println(len(array2))
}
