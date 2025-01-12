package main

import (
	"fmt"
)

func main() {
	fmt.Println(hello("Wilderman"))
	fmt.Println("Suma: ")
	sum, mul := calculate(5, 5)
	fmt.Println("la suma es: ", sum, "la multiplicacion es: ", mul)
}

func hello(name string) string {
	return "Hello: " + name
}

func calculate(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
