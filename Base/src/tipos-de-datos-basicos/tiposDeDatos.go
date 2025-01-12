package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Tipos de datos básicos")

	var entero int = 10
	var enteroPositivo uint = 10

	fmt.Println(enteroPositivo)
	fmt.Println(entero)
	fmt.Println("Valor máximo de un entero", math.MaxInt64)

	fullName := "Juan Perez \t(alias \"El loco\")\n"

	fmt.Println(fullName)

	var a byte = 'A' // byte es un alias para uint8

	s := "hola"

	fmt.Println(a)
	fmt.Println(s[0])

	var r rune = '♥' // rune es un alias para int32

	fmt.Println(r)

	//Variables default
	var (
		enteroDefault  int
		floatDefault   float64
		booleanDefault bool
		stringDefault  string
	)

	fmt.Println(enteroDefault, floatDefault, stringDefault, booleanDefault)

	// Conversionesde datos
	fmt.Println("Conversiones de datos ------------------------------------")
	var entero16 int16 = 10
	var entero32 int32 = 10

	fmt.Println(int32(entero16) + entero32)

	h := "10"
	// Convertir un string a entero
	i, _ := strconv.Atoi(h)

	fmt.Println(i)

	// Convertir un entero a string
	j := strconv.Itoa(1000)

	fmt.Println(j)
}
