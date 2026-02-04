package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Funcion donde los workers elevan al cuadrado los numeros
func cuadradoWorker(id int, nums <-chan int, filtrarPares chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range nums {

		result := num * num
		fmt.Printf("Processor - %d: %d = %d \n", id, num, result)
		filtrarPares <- result
		sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
		time.Sleep(sleepTime)
	}
}

// Funcion que mira si los cuadrados de los numeros elevados en la anterior funcion son pares o impares
func ParesWorker(id int, filtrarPares <-chan int, resultadosPares chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for pares := range filtrarPares {
		if pares%2 == 0 {
			fmt.Printf("Validator - %d: %d is even (passed) \n", id, pares)
			sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
			time.Sleep(sleepTime)
			resultadosPares <- pares

		} else {
			fmt.Printf("Validator - %d: %d is odd (filtered) \n", id, pares)
			sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
			time.Sleep(sleepTime)
		}
	}
}

func main() {
	//Creamos los WaitGroup
	var wgCuadrado sync.WaitGroup
	var wgPares sync.WaitGroup

	//Lista y Canales
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numschan := make(chan int, len(nums))
	paresChan := make(chan int, len(nums))
	resultadosChan := make(chan int)

	//Numero de workers a usar en cada funcion
	workersSquare := 3
	workersPares := 2

	//Lanzamos la funcion cuadradoWorker
	for i := 1; i <= workersSquare; i++ {
		wgCuadrado.Add(1)
		go cuadradoWorker(i, numschan, paresChan, &wgCuadrado)
	}

	//Lanzamos la funcion ParesWorker
	for i := 1; i <= workersPares; i++ {
		wgPares.Add(1)
		go ParesWorker(i, paresChan, resultadosChan, &wgPares)
	}

	//Recogemos los resultados de wait
	resultados := []int{}
	var mx sync.Mutex
	go func() {
		for r := range resultadosChan {
			mx.Lock()
			resultados = append(resultados, r)
			mx.Unlock()
		}
	}()

	//Pasamos los numeros de la lista por el canal
	for _, num := range nums {
		numschan <- num
		fmt.Printf("Generator: produced %d \n", num)
		sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
		time.Sleep(sleepTime)
	}

	//Cerramos canales y le decimos a los WaitGroup que esperen
	close(numschan)
	wgCuadrado.Wait()
	close(paresChan)
	wgPares.Wait()
	close(resultadosChan)

	fmt.Println("Final Results: ", resultados)
}
