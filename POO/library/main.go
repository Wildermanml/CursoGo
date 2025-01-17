package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//Objeto tipo Book
	var myBook = book.NewBook("Moby Dick", "Herman Melville", 300)

	//myBook.PrintInfo()

	myBook.SetTittle("Moby Dick (Edición especial)")
	fmt.Println(myBook.GetTittle())

	var myTextbook = book.NewTextBook("Moby Dick", "Herman Melville", 300, "Mundo", "Primaria")

	//myTextbook.PrintInfo()
	book.Print(myBook)
	book.Print(myTextbook)

	//Objeto tipo animal

	//miPerro := animal.Perro{Nombre: "Negro"}
	//miGato := animal.Gato{Nombre: "Pelusa"}

	//animal.HacerSonido(&miPerro)
	//animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Max"},
		&animal.Perro{Nombre: "Luna"},
		&animal.Gato{Nombre: "Paloma"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
