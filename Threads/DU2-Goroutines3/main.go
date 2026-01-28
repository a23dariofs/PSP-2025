package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DescargarArchivo(id int, extension string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Descarga Inicial: archivo%d.%s \n", id, extension)
	tiempo := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(tiempo)

	tamañoDescarga := rand.Intn(91) + 10
	fmt.Printf("Descarga Completada: file%d.%s (%d MB) \n", id, extension, tamañoDescarga)
}

func main() {
	nombres := []string{"zip", "mp4", "pdf", "doc", "jpg"}
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go DescargarArchivo(i, nombres[i-1], &wg)
	}

	wg.Wait()

	fmt.Println("¡Todas las descargas completadas!")
}
