package main

import "fmt"

func divide(dividendo, divisor int) int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado:", r)
		}
	}()
	// Si el divisor es 0, se lanza un panic
	validateZero(divisor)
	return dividendo / divisor
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por 0")
	}
	fmt.Println("Divisor válido")
}

func main() {
	// Panic se utiliza para errores fatales, como un acceso a un índice fuera de rango, un nil pointer, etc.
	// Se utiliza para interrumpir la ejecución del programa.
	divide(10, 0)
	fmt.Println(divide(10, 2))
}
