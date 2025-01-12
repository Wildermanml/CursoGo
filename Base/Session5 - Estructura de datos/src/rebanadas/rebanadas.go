package main

import (
	"fmt"
)

func main() {
	fmt.Println("Rebanadas")
	// una rebanada es un apuntador a un arreglo y tiene un tamaño y una capacidad (capacidad = tamaño del arreglo).
	diasSemana := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}
	diaRebanada := diasSemana[1:3]
	fmt.Println(diaRebanada)
	fmt.Println("Tamaño de la rebanada", len(diaRebanada))
	fmt.Println("Capacidad de la rebanada", cap(diaRebanada))

	diaRebanada = append(diaRebanada, "Jueves")
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println("funcion make")
	// make es una función que crea una rebanada con un tamaño y una capacidad
	nombres := make([]string, 5, 10)
	fmt.Println("rebanada nombres", nombres)

	//copy hace una copia de una rebanada
	nombres = append(nombres, "Juan", "Pedro", "Maria", "Jose")
	nombres2 := make([]string, len(nombres))
	fmt.Println("rebanada nombres2", nombres2)
	copy(nombres2, nombres)
	fmt.Println(len(nombres2))
	fmt.Println(nombres)

}
