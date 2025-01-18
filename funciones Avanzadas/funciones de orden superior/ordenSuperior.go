package main

import "fmt"

//se caracteriza por ser una funcion que devuelve otra funcion.

func double(f func(int) int, x int) int {

	return f(x * 2)

}

func addOne(x int) int {
	return x + 1
}

func main() {
	r := double(addOne, 3)
	fmt.Println("Resultado: ", r)
}
