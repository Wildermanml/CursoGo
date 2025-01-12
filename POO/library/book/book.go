package book

import "fmt"

//Comparando con programacion orientada a objetos, este es un modelo
type Book struct {
	//Encapsulamiento: los nombres de las variables, si inician con min es privada, con mayus es publica o exportable, ejemplo Tittle: Publica, tittle: Privada(solo se usa en la misma clase)
	tittle string
	author string
	pages  int
}

//Este un metodo
func (b *Book) PrintInfo() {
	fmt.Printf("Tittle: %s\nAuthor:  %s\nPages: %d\n", b.tittle, b.author, b.pages)
}

// Simula un Constructor
func NewBook(tittle string, author string, pages int) *Book {
	return &Book{
		author: author,
		pages:  pages,
		tittle: tittle,
	}
}

//Getters
func (b *Book) SetTittle(tittle string) {
	b.tittle = tittle
}

//Setters
func (b *Book) GetTittle() string {
	return b.tittle
}
