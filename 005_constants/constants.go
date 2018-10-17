package main

import (
	"fmt"
)

//const applicationName = "App Const"
// const (
//_ = iota
//B = iota
//C = iota
//D = iota + 5
// )

func main() {
	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	//fmt.Println(D)
	//const number = 100
	//fmt.Println(applicationName)
	//fmt.Println(number)

	//var numberTwo = 10
	//result := number + numberTwo
	//fmt.Println(result)

	x, y, z := TypesGo()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func TypesGo() (int, bool, string) {

	return 100, false, "string"

}
