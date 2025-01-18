package main

import "fmt"

// funciones variadicas: son funciones que pueden recibir n parametros
func suma(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	return total

}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v", dato, dato)
	}
}

func main() {
	//fmt.Println(suma(1, 2, 3, 4, 5, 6))
	imprimirDatos("Hola", true, 10.4)
}
