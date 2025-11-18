package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Longitud de los argumentos no válida")
		return
	}

	strNum := os.Args[1]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return
	}
	file.Close()

	file, err = os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo en formato append: ", err)
		return
	}
	defer file.Close()

	WriteNumbers(num, file)

}

func WriteNumbers(num int, file *os.File) {
	if num == 0 {
		time.Sleep(500)
	} else {
		time.Sleep(time.Duration(num) * time.Second)
	}

	fmt.Println(num)
	file.WriteString(strconv.Itoa(num) + " \n")
}
