package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Switch")
	os := runtime.GOOS
	fmt.Println("Sistema operativo: ", os)

	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("No se pudo determinar el sistema operativo")
	}

}
