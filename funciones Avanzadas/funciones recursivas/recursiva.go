package main

import "fmt"

//Funcion recursiva, es una funcion que se llama asi misma

func factorial(n int) int {
	fmt.Print(n)
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(10))
}
