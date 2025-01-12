package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Proyecto")
	fmt.Println("Numero aleatorio: ", rand.Intn(100))
	play()

}

func play() {
	fmt.Println("Jugando")
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 7
	for intentos <= maxIntentos {
		fmt.Println("Ingrese un numero: ")
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("Ganaste!")
			return // puede agregarse un break para terminar el for, pero el return indica que ya termino la funcion play
		} else if numIngresado > numAleatorio {
			fmt.Println("Numero aleatorio es menor")
		} else if numIngresado < numAleatorio {
			fmt.Println("Numero aleatorio es mayor")
		}
		fmt.Println("Intentos : ", intentos)
		intentos++
	}
	fmt.Println("Se acabaron los intentos! El numero era: ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	fmt.Println("Queres Jugar de nuevo? (s/n)")
	var respuesta string
	fmt.Scanln(&respuesta)
	if respuesta == "s" {
		play()
	} else if respuesta == "n" {
		fmt.Println("Gracias por jugar!")
	} else {
		fmt.Println("Respuesta invalida, debe responder s o n")
		jugarNuevamente()
	}
}
