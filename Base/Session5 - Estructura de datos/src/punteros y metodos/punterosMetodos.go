package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main() {
	fmt.Println("Punteros y métodos")
	var x int = 10
	// este apunta a la dirección de memoria de x
	fmt.Println("Valor de x: ", x)
	editar(&x)
	fmt.Println("Valor de x modificado por editar() ", x)

	p := persona{nombre: "Juan", edad: 20}
	p.diHola()

}

func editar(x *int) {
	*x = 50
}

func (p *persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)

}
