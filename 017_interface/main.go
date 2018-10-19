package main

import (
	"fmt"
)

// Полиморфизм - это свойство классов иметь одинаковые методы, которые работают по разному в контексте объектов

type Document struct {
	Date       string
	Number     string
	NumOfPages int
}

type PersonCard struct {
	Date      string
	FirstName string
	LastName  string
	Age       int
}

type PrintInterface interface {
	CheckData()
}

func (d *Document) CheckData() {
	if d.Date != "" {
		fmt.Println("Date in the doc - correct")
	} else {
		fmt.Println("Date in the doc is incorrect")
	}
}

func (p *PersonCard) CheckData() {
	if p.Date != "" {
		fmt.Println("Date in the person card - correct")
	} else {
		fmt.Println("Date in the person card is incorrect")
	}
}

func main() {

	doc := new(Document)
	doc.Date = "1.10.2018"
	doc.NumOfPages = 3
	doc.Number = "A - 100"

	doc2 := new(Document)

	pCard := new(PersonCard)
	pCard.Date = "1/10/2019"
	pCard.Age = 21
	pCard.FirstName = "user"
	pCard.LastName = "lastname"

	sl := []PrintInterface{doc, pCard, doc2}

	PrintAnyDoc(sl)

}

func PrintAnyDoc(anyDoc []PrintInterface) {
	for _, v := range anyDoc {
		fmt.Println(v)
	}
}
