package main

import (
	"log"
	"os"
)

func main() {
	//Registro de errores
	file, error := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetPrefix("registroError.main():")

	if error != nil {
		log.Fatal("Error al abrir el archivo de registro")
	}
	log.SetOutput(file)
	log.Println("Oye, soy un log")

	//log.Fatal("Error")
	//log.Panic("Mensaje de registro de Panic, termina la ejecución del programa, devuelve pila de ejecución")
	//log.Println("Mensaje de registro de errores con salto de linea")

}
