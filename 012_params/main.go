package main

import "fmt"

func main() {

	sliceOfInt := [][]int string{{33, 23, 22}, {2, 3, 4}}
	value := calculate(sliceOfInt...)
	fmt.Println(value)
}

func calculate(numbers ...[]int) int {

	var result int

	//fmt.Printf("%T", numbers)

	for _, number := range numbers {
		for _, number2 := range number {
			result += number2
		}

	}

	return result

}
