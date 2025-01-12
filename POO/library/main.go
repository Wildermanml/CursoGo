package main

import (
	"fmt"
	"library/book"
)

func main() {
	//Objeto tipo Book
	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	myBook.PrintInfo()

	myBook.SetTittle("Moby Dick (Edición especial)")
	fmt.Println(myBook.GetTittle())
}
