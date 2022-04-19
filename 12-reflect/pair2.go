package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (b *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b : pair<type:Book, value:book{}地址>
	b := &Book{}

	//r: pair<type:, value:>
	var r Reader

	//r: pair<type:Book, value:book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	//r: pair<type:Book, value:book{}地址>
	w = r.(Writer)
	w.WriteBook()
}
