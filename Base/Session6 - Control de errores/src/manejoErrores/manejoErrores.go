package main

import (
	"errors"
	"fmt"
)

func divide(diviendo, divisor int) (int, error) {
	if divisor == 0 {
		// return 0, fmt.Errorf("No se puede dividir por cero")
		return 0, errors.New("No se puede dividir por cero")
	}
	return diviendo / divisor, nil
}

func main() {
	result, err := divide(10, -1)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Resultado: ", result)
}
