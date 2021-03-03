package main

import "fmt"

func main() {
	type Books struct {
		title  string
		year   int
		author string
		page   int
	}
	var books1 Books
	books1.title = "Go Learn"
	books1.year = 18
	books1.author = "guoliang"
	books1.page = 999
	fmt.Println(books1)
	fmt.Println(len(books1))
}
