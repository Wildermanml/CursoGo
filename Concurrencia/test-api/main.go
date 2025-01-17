package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//canales
	///////crear canal
	//canal := make(chan int)
	////// Enviar datos al canal
	//canal <- 15

	///Recibir datos del canal
	//valor := <-canal

	ch := make(chan string)
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		//Sin concurrencia
		//checkApi(api)

		//Con concurrencia
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {

		fmt.Println(<-ch)
	}

	//time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())

}

func checkApi(api string, ch chan string) {

	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error ¡%s está caido!\n", api)
		return
	}

	ch <- fmt.Sprintf("Success!! %s\n", api)
}
