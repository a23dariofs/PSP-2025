package main

import (
	"fmt"
	"os"
)

func main() {
	//Lee que se hayan pasado todos los argumentos y si no termina el programa (son dos porque el args[0] es el nombre del programa y args[1] el argumento en string)
	if len(os.Args) < 2 {
		return
	}

	//Añadimos a una variable llamada numStr el argumento pasado en el padre y lo imprimimos
	numStr := os.Args[1]

	fmt.Println(numStr)

	//Creamos el file
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	//Cierra el file
	defer file.Close()

	//Esto es para escribir en el file cada uno de los numeros con un salto de linea
	file.WriteString(numStr + "\n")
}
