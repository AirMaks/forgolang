package main

import "fmt"

func main() {
	var a int = 5
	var b, c string = "b", "c"
	var d bool = true

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
	fmt.Printf("%T - %v\n", d, d)

	e := 100
	f := "f"
	g := true

	fmt.Printf("%T - %v\n", e, e)
	fmt.Printf("%T - %v\n", f, f)
	fmt.Printf("%T - %v\n", g, g)

	var zero int
	var emptyString string
	var boolFalse bool

	fmt.Printf("%T - %v\n", zero, zero)
	fmt.Printf("%T - %v\n", emptyString, emptyString)
	fmt.Printf("%T - %v\n", boolFalse, boolFalse)

}
