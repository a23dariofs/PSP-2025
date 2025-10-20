package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("El funcionamiento del programa es: <program> -n 3 1 2 3...")
		return
	}

	casos, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: numCasos debe ser un número")
		return
	}

	if len(os.Args) != 2+casos*2 {
		fmt.Println("Error: se deben proporcionar numEscalones y escalonesMax para cada caso")
		return
	}

	for i := 0; i < casos; i++ {
		numEscalones, err1 := strconv.Atoi(os.Args[2+i*2])
		escalonesMax, err2 := strconv.Atoi(os.Args[3+i*2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error: numEscalones y escalonesMax deben ser números")
			return
		}

		fmt.Printf("Caso %d: El número mínimo de saltos necesarios es: %d\n",
			i+1, saltosNecesarios(numEscalones, escalonesMax))
	}
}

func saltosNecesarios(escalonesTotales int, escalonesXsalto int) int {
	resultado := escalonesTotales / escalonesXsalto
	if escalonesTotales%escalonesXsalto != 0 {
		resultado++
	}
	return resultado
}
