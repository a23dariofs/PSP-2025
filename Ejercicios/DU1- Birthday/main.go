package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Introduzca el numero de casos de prueba: ")
	var casos int
	fmt.Scan(&casos)

	for i := 0; i < casos; i++ {
		count := 0
		fmt.Print("Introduzca la edad: ")
		var edad int64
		fmt.Scan(&edad)

		edadBinario := strconv.FormatInt(edad, 2)

		for _, c := range edadBinario {
			if c == '1' {
				count++
			}
		}
		fmt.Println("Velas a encender:", count)
	}
}
