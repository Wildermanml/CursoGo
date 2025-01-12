package main

import (
	"fmt"
	"log"

	"github.com/wildermanml/greetings"
)

func main() {
	log.SetPrefix("greetings - ")

	names := []string{"Wilderman", "Gopher", "Luisa"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range message {
		fmt.Println(message)
	}

}
