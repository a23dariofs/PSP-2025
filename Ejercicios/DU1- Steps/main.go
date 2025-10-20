package main

import (
	"fmt"
)

func main() {
	fmt.Print("Introduzca el numero de casos de prueba: ")
	var casos int
	fmt.Scan(&casos)

	for i := 0; i < casos; i++ {
		fmt.Println("Intruduzca en una linea el numero de escalones y cuantos escalones se pueden subir de una vez: ")
		var numEscalones int
		var EscalonesMax int
		fmt.Scan(&numEscalones, &EscalonesMax)
		fmt.Println("El numero minimo de saltos necesarios es: ", saltosNecesarios(numEscalones, EscalonesMax))
	}
}

func saltosNecesarios(escalonesTotales int, escalonesXsalto int) int {
	resultado := escalonesTotales / escalonesXsalto
	if escalonesTotales%escalonesXsalto != 0 {
		resultado++
	}
	return resultado
}
