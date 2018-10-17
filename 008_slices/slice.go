package main

import (
	"fmt"
)

func main() {
	//var sliceOne []int
	//fmt.Println(len(sliceOne))
	//fmt.Println(cap(sliceOne))

	//sliceOne = append(sliceOne, 10)
	//sliceOne = append(sliceOne, 20)
	//sliceOne = append(sliceOne, 30)

	//fmt.Println(len(sliceOne))
	//fmt.Println(cap(sliceOne))

	//fmt.Println(sliceOne[:3])

	// slice := make([]int, 5, 20)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	mySlice := []int{1, 2, 5}
	fmt.Println(mySlice[2:])
}
