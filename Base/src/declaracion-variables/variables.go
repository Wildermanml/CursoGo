package main

import (
	"fmt"
)

// Declaración de constantes
const PI float32 = 3.1416
const (
	x = 100
	y = 0b1010 // 10 en binario
	z = 0o12   // 10 en octal
	w = 0x1A   // 10 en hexadecimal
)

const (
	Domingo = iota + 1 // iota = 0 y se incrementa en 1 en cada constante, asi lunes tendra valor 1 martes 2 y asi sucesivamente.
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Declaración de variables")
	// Declaración de variables
	var firstName string
	firstName = "Juan"
	fmt.Println(firstName)

	// Declaración de variables con inferencia de tipo
	lastName := "Perez"
	fmt.Println(lastName)

	// Declaración de múltiples variables
	var age, height int
	age = 30
	height = 180
	fmt.Println(age, height)

	// Declaración de multiples variables
	var (
		weight   int    = 80
		eyeColor string = "brown"
	)

	fmt.Println(firstName, weight, eyeColor)

	// Declaración de variables sin tipo de dato
	var (
		companyName = "Google"
		companySize = 100000
	)

	fmt.Println(companyName, companySize)

	// Declaración de variables con inicialización de valores
	var nombre, apellido, edad = "Juan", "Perez", 20

	fmt.Println(nombre, apellido, edad)

}
