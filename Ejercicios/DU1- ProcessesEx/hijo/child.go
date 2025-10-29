package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Debes pasar un número como argumento")
		return
	}

	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("El argumento debe ser un número entero")
		return
	}

	// Abrir (o crear) el archivo y vaciarlo
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo:", err)
		return
	}
	file.Close() // cerramos tras truncar

	// Reabrir en modo append
	file, err = os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo en modo append:", err)
		return
	}
	defer file.Close()

	writeNumbers(num, file)
}

func writeNumbers(num int, file *os.File) {
	//Porque al limpiar el archivo sino le pones un tiempo de espera al primer numero se lo come también
	if num == 0 {
		time.Sleep(time.Duration(500) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(num) * time.Second)
	}
	fmt.Println(num)
	file.WriteString(strconv.Itoa(num) + "\n")

}
