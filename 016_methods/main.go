package main

import (
	"fmt"
)

type EDocument struct {
	Number        string
	Date          string
	NumberOfPages int
	// Footer - анонимное поле сразу определенное типом структуры
	FooterDoc Footer // именованное поле с типом

}

type Footer struct {
	Author string
}

func (doc EDocument) showAuthor() {
	fmt.Println(doc.FooterDoc.Author)
}

func (doc *EDocument) showNumberOfPages() {
	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument) ClearAllPages() {
	doc.NumberOfPages = 0
	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument) SetPrefixAuthor() {
	doc.FooterDoc.Author = "The author is " + doc.FooterDoc.Author
	fmt.Println(doc.FooterDoc.Author)
}

func main() {

	doc1 := EDocument{
		Number:        "001-A",
		Date:          "10.09.2018",
		NumberOfPages: 2,
		FooterDoc: Footer{
			Author: "John Galdman",
		},
	}

	var doc2 EDocument
	doc2.Number = "001-A"
	doc2.Date = "10.09.2018"
	doc2.NumberOfPages = 72
	doc2.FooterDoc.Author = "James Cameron"

	doc3 := new(EDocument)
	doc3.Number = "003-A"
	doc3.Date = "23/092018"
	doc3.NumberOfPages = 7
	doc3.FooterDoc.Author = "No author"

	// fmt.Printf("%T - %v\n", doc1, doc1)
	// fmt.Printf("%T - %v\n", doc2, doc2)
	// fmt.Printf("%T - %v\n", doc2.FooterDoc, doc2.FooterDoc)
	// fmt.Printf("%T - %v", doc3, doc3)

	doc1.SetPrefixAuthor()
	doc1.ClearAllPages()
}
