package main

import (
	"fmt"
)

func main() {
	var simpleMap = map[int]string{}

	simpleMap[198] = "some word"
	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)
	anotherMap["one"] = 198
	anotherMap["two"] = 298
	anotherMap["three"] = 398
	fmt.Println(anotherMap)

	delete(anotherMap, "three")
	fmt.Println(anotherMap)

	slice := []int{1, 2, 3, 4}
	for _, v := range slice {
		fmt.Println(v)
		if v == 2 {
			fmt.Println("ok")
		}
	}
}
