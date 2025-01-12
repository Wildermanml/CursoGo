package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("If Else")
	fmt.Println("Juego de adivinar un numero")

	t := time.Now()
	fmt.Println(t.Hour())

	fmt.Println("Tienes que adivinar el numero secreto")

	var number int
	fmt.Println("Ingresa un numero")
	fmt.Scanln(&number)

	if number == t.Hour() {
		fmt.Println("Adivinaste el numero")
	} else if number == t.Hour()-1 || number == t.Hour()+1 {
		fmt.Println("Casi adivinas el numero")
	} else {
		fmt.Println("No adivinaste el numero")

	}
}
