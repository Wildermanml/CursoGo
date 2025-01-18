package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	//PrintList("Juan", 17, "Guarne")
	//PrintListAny("Victor", 35, "Mercadolibre", 104.5)
	fmt.Println(SumConstraints[uint](10, 11, 23))
	fmt.Println(SumConstraints[float32](10.0, 11.8, 23.36845, 45.4))

	//llamado a includes
	strings := []string{"a", "b", "c", "d", "e"}
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("===========LLamado a include============")
	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(numbers, 20))

	fmt.Println("========LLamado funcion Filter===========")
	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	fmt.Println("========Uso de estructura generica===========")
	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"2", "Zapatos", 50}
	fmt.Println("Product 1", product1)
	fmt.Println("Product 2", product2)

}

// funcion generica sin usar any
func PrintList(List ...interface{}) {
	for _, value := range List {
		fmt.Println(value)
	}
}

// funcion generica usando any
func PrintListAny(List ...any) {
	for _, value := range List {
		fmt.Println(value)
	}
}

// Restricciones
type integer int

// type de Numbers permitidos para la suma
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// Go ya tiene paquetes de restricciones o constraints predefinidos, estos se pueden encontrar en su documentacion
//https://pkg.go.dev/golang.org/x/exp/constraints

func Sum[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func SumConstraints[T constraints.Float | constraints.Integer](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

// Estructura generica, si observamos, Id es tipo T que a su vez esta entre unit o string
type Product[T uint | string] struct {
	Id          T
	Descripcion string
	Precio      float32
}
