package main

import (
	"fmt"
	"github.com/AirMaks/forgolang/002_package/nested/hello"
	"github.com/AirMaks/forgolang/002_package/nested/say"
)

func main() {

	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello() + hello.NotVisible)

}
