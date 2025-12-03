package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const totalDias = 3650
	const partes = 10

	// Crear y rellenar el array de temperaturas
	temperaturas := make([]int, totalDias)

	for i := 0; i < totalDias; i++ {
		temperaturas[i] = rand.Intn(81) - 30 // -30 a 50
	}

	// Canal para recoger las sumas parciales
	resultados := make(chan int)

	// Tamaño de cada trozo
	tam := totalDias / partes

	// Lanzar goroutines
	for i := 0; i < partes; i++ {
		inicio := i * tam
		fin := inicio + tam

		if i == partes-1 {
			fin = totalDias
		}

		go func(id, start, end int) {
			suma := 0
			for j := start; j < end; j++ {
				suma += temperaturas[j]
			}
			fmt.Printf("Goroutine %d suma parcial: %d\n", id, suma)
			resultados <- suma
		}(i, inicio, fin)
	}
	// Recoger sumas parciales
	sumaTotal := 0
	for i := 0; i < partes; i++ {
		sumaTotal += <-resultados 
	}

	// Calcular media
	media := float64(sumaTotal) / float64(totalDias)

	fmt.Println("--------------------------------------")
	fmt.Printf("Suma total: %d\n", sumaTotal)
	fmt.Printf("Temperatura media: %.2f°C\n", media)
}
