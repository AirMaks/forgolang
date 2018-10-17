package main

import (
	"fmt"
)

var car = "BMW"

func main() {

	fmt.Println(car)
	checkSpareWheel()

}

func checkSpareWheel() {
	if car == "BMW" {
		fmt.Println("A spare wheel exists.")
	}
}
