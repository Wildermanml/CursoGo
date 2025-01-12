package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Contact struct {
	Nombre   string `json:"name"`
	Telefono string `json:"phone"`
	Email    string `json:"email"`
}

func saveContacToFile(contactos []Contact) error {
	//Creación de un archivo
	file, err := os.Create("contacts.json")
	defer file.Close()

	if err != nil {
		log.Fatal("err al abrir el archivo de registro")
		return err
	}

	//Codificación de los contactos en formato JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contactos)

	if err != nil {
		log.Fatal("err al codificar los contactos en formato JSON")
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	//Apertura de un archivo
	file, err := os.Open("contacts.json")
	defer file.Close()
	if err != nil {
		fmt.Println("err al abrir el archivo de registro")
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		log.Fatal("err al decodificar los contactos en formato JSON")
		return err
	}
	return nil

}

func main() {
	//Creación de un slice de Contacto
	var contacts []Contact

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("err al cargar los contactos")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("====Gestión de contactos======\n",
			"1. Crear contacto\n",
			"2. Listar contactos\n",
			"3. Salir\n",
			"Opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			log.Fatal("err al leer la opción")
		}

		switch opcion {

		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Telefono, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			contacts = append(contacts, c)

			if err := saveContacToFile(contacts); err != nil {
				log.Fatal("err al guardar los contactos")
			}

		case 2:
			fmt.Println("Lista de contactos")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s - Telefono: %s - Email: %s\n", index+1, contact.Nombre, contact.Telefono, contact.Email)
			}
			fmt.Println("=====================================")
		case 3:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opción no válida")
		}

	}
}
