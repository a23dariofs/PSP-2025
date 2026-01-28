package main

import (
	"fmt"
	"sync"
)

func sumNumbers(id int, numeroInicio int, numeroFinal int) int {
	resultado := 0

	for i := numeroInicio; i <= numeroFinal; i++ {
		resultado += i
	}

	fmt.Printf("Trabajador %d: Suma de %d a %d = %d \n", id, numeroInicio, numeroFinal, resultado)

	return resultado
}

func main() {
	var wg sync.WaitGroup
	totalSum := 0
	nums := 1
	mutex := sync.Mutex{}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(nums int) {
			defer wg.Done()
			mutex.Lock()
			totalSum += sumNumbers(i, nums, nums+24)
			mutex.Unlock()
		}(nums)
		nums += 25
	}
	wg.Wait()

	fmt.Printf("Suma total: %d", totalSum)

}
