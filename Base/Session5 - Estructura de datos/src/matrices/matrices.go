package main

import (
	"fmt"
)

func main() {
	fmt.Println("Matrices")
	var matriz [3]int
	matriz[0] = 1
	matriz[1] = 2
	matriz[2] = 3
	fmt.Println(matriz)

	matriz2 := [3]int{1, 2, 3}
	fmt.Println(matriz2)

	matriz3 := [...]int{1, 2, 3}
	fmt.Println(matriz3)

	for index, value := range matriz3 {
		fmt.Println(index, value)
	}

	fmt.Println("Matriz bidimensional")
	var matrizBi [2][2]int
	matrizBi[0][0] = 1
	matrizBi[0][1] = 2
	matrizBi[1][0] = 3
	matrizBi[1][1] = 4
	fmt.Println(matrizBi)
}
