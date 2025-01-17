package animal

import "fmt"

type Animal interface {
	Sonido()
}

// estructura perro y sus metodos

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre, " Hace Grrrrr!!!")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre, " Hace Miauuuu!")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
