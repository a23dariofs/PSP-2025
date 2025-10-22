package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	numStr := os.Args[1]

	fmt.Println(numStr)

	var flag int
	if numStr == "0" {
		flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	} else {
		flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	}

	file, err := os.OpenFile("output.txt", flag, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Escribimos directamente el string (sin fmt)
	file.WriteString(numStr + "\n")
}
