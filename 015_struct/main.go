package main

import "fmt"

type EDocument struct {
	Number        string
	Date          string
	NumberOfPages int
	Footer
}

type Footer struct {
	Author string
}

func main() {

	doc1 := EDocument{
		Number:        "001-A",
		Date:          "10.09.2018",
		NumberOfPages: 10,
		Footer: Footer{
			Author: "David Becks",
		},
	}

	var doc2 EDocument
	doc2.Number = "002-A"
	doc2.Date = "10.09.2018"
	doc2.NumberOfPages = 4
	doc2.Footer.Author = "Steven Spilberg"

	fmt.Printf("%T - %v\n", doc1, doc1)
	fmt.Printf("%T - %v\n", doc2, doc2)
}
