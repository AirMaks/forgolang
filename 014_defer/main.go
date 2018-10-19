package main

import "fmt"

func main() {

	defer one()
	defer two()
	defer three()
	four()

}

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}

func four() {
	fmt.Println("four")
}
