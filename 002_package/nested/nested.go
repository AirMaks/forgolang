package main

import (
	"fmt"
	"github.com/forgolang/002_package/nested/hello"
	"github.com/forgolang/002_package/nested/say"
)

func main() {

	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())

}
