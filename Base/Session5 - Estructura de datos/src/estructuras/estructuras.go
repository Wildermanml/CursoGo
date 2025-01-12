package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
	correo string
}

func main() {
	fmt.Println("Estructuras de datos")
	var p Persona
	p.Nombre = "Wilderman"
	p.Edad = 35
	p.correo = "Wilderman@gmail.com"

	p2 := Persona{Nombre: "juan", Edad: 35, correo: "juan@gmail.com"}

	p2.Nombre = "Pedro"
	fmt.Println(p, p2)

}
