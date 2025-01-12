package main

import (
	"fmt"
	"os"
)

func main() {
	//defer es para ejecutar una función al final de la función actual
	defer fmt.Println("Esto se ejecuta al final de la función")
	fmt.Println("Esto se ejecuta primero")
	fmt.Println("Esto se ejecuta después")

	file, error := os.Create("hola.txt")
	if error != nil {
		fmt.Println("Error al crear el archivo")
		return
	}

	_, err := file.Write([]byte("Hola mundo"))
	if err != nil {
		fmt.Println("Error al escribir en el archivo")
		return
	}

	defer file.Close()
}
