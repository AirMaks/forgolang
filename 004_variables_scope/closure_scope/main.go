package main

import "fmt"

func main() {
	x := 100
	fmt.Println(x)
	someFunc()
	someFunc2()
	fmt.Println(miniGlobal)
	dialog()
}

func someFunc() {
	x := 700
	fmt.Println(x)

}

func someFunc2() {
	var test = "Okay"
	var point = 4.4

	fmt.Println(test)
	fmt.Println(point)
}

var miniGlobal = true

func dialog() {
	var x = "Hello!"
	var y = "Hi:)"

	fmt.Println(x)
	{
		fmt.Println(y)
		var z = "How are you?"
		fmt.Println(z)
	}

}
