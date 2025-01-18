package main

import "fmt"

//Un closure en Go es una función que "recuerda" las variables del entorno donde fue creada, incluso si esa función se ejecuta fuera de ese entorno.
//Esto significa que una función puede acceder y usar variables definidas fuera de ella, como si las "cerrara" dentro de su alcance.

func incrementar() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incrementar()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(nextInt())

	fmt.Println(nextInt())

}
