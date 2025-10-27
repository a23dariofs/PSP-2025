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

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo:", err)
		return
	}
	file.Close() // cerramos después de truncar

	// Abrimos el archivo para append
	file, err = os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo en modo append:", err)
		return
	}
	defer file.Close()

	num, _ := strconv.Atoi(numStr)
	writeNumbers(num, file)
}

func writeNumbers(num int, file *os.File) {
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println(num)
	file.WriteString(strconv.Itoa(num) + "\n")
}
