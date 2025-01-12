package main

import (
	"fmt"
)

func main() {
	fmt.Println("For")
	fmt.Println("Ciclo For")
	fmt.Println("Imprimir los numeros del 1 al 10")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Imprimir los numeros del 10 al 1")

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println("Imprimir los numeros pares del 1 al 10")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Imprimir los numeros impares del 1 al 10")

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Imprimir los numeros del 1 al 10 con un ciclo infinito")

	i := 1
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}

	fmt.Println("Imprimir los numeros del 1 al 10 con un ciclo infinito y saltando el numero 5")

	i = 1
	for {
		if i == 5 {
			i++
			continue //Continue se usa para saltar a la siguiente iteracion, es decir no imprime el numero 5.
		}
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}
}
