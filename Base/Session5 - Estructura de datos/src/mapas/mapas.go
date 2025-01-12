package main

import (
	"fmt"
)

func main() {
	fmt.Println("Mapas")
	// Un mapa es una estructura de datos que almacena pares clave-valor
	// Se declara con la palabra reservada map seguida de la clave y el tipo de valor
	// Se inicializa con la función make
	// Se accede a un valor con la clave
	// Se añade un valor con la clave
	// Se elimina un valor con la función delete
	// Se recorre con un ciclo for
	// Se puede hacer una asignación múltiple

	colors := map[string]string{
		"rojo":  "#ff0000",
		"verde": "#00ff00",
		"azul":  "#0000ff",
	}

	fmt.Println(colors)
	fmt.Println(colors["rojo"])

	// Se añade un valor con la clave
	colors["blanco"] = "#ffffff"
	fmt.Println(colors)

	// validar si una clave existe, si esta validacion no se agrega el valor es un valor vacio.
	valorRojo, verficador := colors["rojo"]
	if verficador {
		fmt.Println(valorRojo, verficador)
	} else {
		fmt.Println("La clave no existe")
	}

	//Eliminar un valor
	delete(colors, "rojo")
	fmt.Println(colors)

	//Recorrer un mapa
	for clave, valor := range colors {
		fmt.Println(clave, valor)
	}
}
