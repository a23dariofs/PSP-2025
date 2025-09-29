package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Debes meter al menos 2 argumentos: <program> -n 3 1 2 3...")
		return
	}

	for _, arg := range os.Args[1:] { //
		edad, error := strconv.ParseInt(arg, 10, 64)
		if error != nil { // nil es lo mismo que null
			fmt.Println("Error: argumento no valido", arg)
			continue
		}

		edadBinario := strconv.FormatInt(edad, 2)

		count := 0

		for _, c := range edadBinario {
			if c == '1' {
				count++
			}
		}
		fmt.Println("Velas a encender:", count)
	}
}
