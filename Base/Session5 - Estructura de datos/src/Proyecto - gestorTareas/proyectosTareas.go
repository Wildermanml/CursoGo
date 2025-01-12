package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *ListaTareas) marcarCompletado(i int) {
	lt.tareas[i].completado = true
}

func (lt *ListaTareas) editarTarea(i int, t Tarea) {
	lt.tareas[i] = t
}

func (lt *ListaTareas) mostrarTareas() {
	fmt.Println("Lista de tareas ###########################################")
	for i, t := range lt.tareas {
		fmt.Printf("%d. %s, estado: %t\n", i+1, t.nombre, t.completado)
	}
}

func (lt *ListaTareas) eliminarTarea(i int) {
	lt.tareas = append(lt.tareas[:i], lt.tareas[i+1:]...)
}

func mostrarMenu() {
	fmt.Println("Seleccione una opción:\n",
		"1. Agregar tarea\n",
		"2. Marcar tarea como completada\n",
		"3. Mostrar tareas\n",
		"4. Editar tarea\n",
		"5. Eliminar tarea\n",
		"6. Salir.\n")
}

func main() {
	fmt.Println("Proyecto - gestor de tareas")

	// Instancia de ListaTareas
	lista := ListaTareas{}

	var op, i int
	var nombre, descripcion string
	leer := bufio.NewReader(os.Stdin)
	for op != 6 {
		mostrarMenu()

		fmt.Print("Opción: ")
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("Ingrese el nombre de la tarea:")
			nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			descripcion, _ = leer.ReadString('\n')
			t := Tarea{nombre, descripcion, false}
			lista.agregarTarea(t)
		case 2:
			fmt.Println("Ingrese el número de la tarea a marcar como completada:")
			fmt.Scanln(&i)
			lista.marcarCompletado(i - 1)
		case 3:
			lista.mostrarTareas()
		case 4:
			fmt.Println("Ingrese el número de la tarea a editar:")
			fmt.Scanln(&i)
			fmt.Println("Ingrese el nombre de la tarea:")
			nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			descripcion, _ = leer.ReadString('\n')
			t := Tarea{nombre, descripcion, false}
			lista.editarTarea(i-1, t)
		case 5:
			fmt.Println("Ingrese el número de la tarea a eliminar:")
			fmt.Scanln(&i)
			lista.eliminarTarea(i - 1)
		case 6:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida")

		}

	}
}
