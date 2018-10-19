package main

import (
	"fmt"
)

func main() {
	sum(5, 5)
	fmt.Println(question("question"))
	result, logic := someFunc(-1, 4, 300)
	fmt.Println(result, logic)

	someFuncTwo(777, 888)

	fmt.Println(onlyTypes(10, "23"))
}

func sum(a int, b int) {

	var z int
	z = a + b
	fmt.Println(z)

}

func question(q string) bool {

	if q == "question" {
		return true
	} else {
		return false
	}

}

func someFunc(a, b, c int) (int, bool) {

	if a > 0 && b > 5 {
		return 10, true
	} else if c > 100 {
		return 100, false
	} else {
		return 0, false
	}

}

func someFuncTwo(z, _ int) {

	fmt.Println(z)

}

func onlyTypes(int, string) bool {
	return true
}
