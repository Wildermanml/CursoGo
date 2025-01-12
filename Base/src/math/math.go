package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Funciones math")

	a := 10
	b := 3

	c := a + b

	fmt.Println(c)

	fmt.Println("Raiz cuadrada de 16: ", math.Sqrt(16))
	fmt.Println("Redondear 3.14: ", math.Round(3.14))
	fmt.Println("modulo de 10/3: ", math.Mod(10, 3))
	fmt.Println("Maximo de 10 y 3: ", math.Max(10, 3))
	fmt.Println("Minimo de 10 y 3: ", math.Min(10, 3))
	fmt.Println("Exponencial de 3: ", math.Exp(3))
	fmt.Println("Logaritmo de 10: ", math.Log(10))
	fmt.Println("Seno de 90: ", math.Sin(90))

}
