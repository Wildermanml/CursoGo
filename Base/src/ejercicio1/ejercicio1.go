package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ejercicio 1")

	var lado1 int
	var lado2 int
	fmt.Println("Ingresa lado 1")
	fmt.Scanln(&lado1)
	fmt.Println("Ingresa lado 2")
	fmt.Scanln(&lado2)

	fmt.Println("Lados del tringulo rectangulo: ", lado1, lado2)
	hipotenusa := math.Sqrt(float64(lado1*lado1 + lado2*lado2))
	fmt.Println("Hipotenusa: ", hipotenusa)

	fmt.Println("Area del triangulo segun teorema de pitagoras: ", float64(lado1*lado2)/2)
	fmt.Println("Perimetro del triangulo: ", float64(lado1+lado2)+hipotenusa)
	fmt.Println("Presicion de 2 decimales en el area: ", fmt.Sprintf("%.2f", float64(lado1*lado2)/2))

}
