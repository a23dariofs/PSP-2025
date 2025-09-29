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

func saltosNecesarios(EscalonesTotales int, EscalonesXsalto int) int {
	var division int = (EscalonesTotales % EscalonesXsalto)
	var resultado int = (EscalonesTotales / EscalonesXsalto)
	if division != 0 {
		return resultado + 1
	} else if division == 0 {
		return resultado
	}
	return 0
}
