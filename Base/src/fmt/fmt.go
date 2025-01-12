package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("funcionalidad para imprimir en consola")
	fmt.Print(quote.Go())
	fmt.Println(("sin salto de linea"))

	var name string
	var lastName string
	var edad int

	fmt.Println("Ingresa tu nombre")
	fmt.Scanln(&name)
	fmt.Println("Ingresa tu apellido")
	fmt.Scanln(&lastName)
	fmt.Println("Ingresa tu edad")
	fmt.Scanln(&edad)

	gretting := fmt.Sprintf("Hola, me llamo %s %s y tengo %d años\n", name, lastName, edad)

	fmt.Println("Usando sprintf", gretting)

}
