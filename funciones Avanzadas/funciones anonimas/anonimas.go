package main

import "fmt"

//funcion anonima, es aquella que no tiene nombre

func main() {
	//definicion funcion anonima
	saludo := func(name string) {
		fmt.Printf("Hola,%s!\n", name)
	}
	saludar("Wilderman ", saludo)

	fmt.Println("------------Funciones matematicas-----------------")

	duplicar := aplicar(duplicar, 15)
	fmt.Printf("El numero duplicado es: %v\n", duplicar)
}

func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}
