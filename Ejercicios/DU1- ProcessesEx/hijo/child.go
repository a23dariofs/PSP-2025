package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	file.Close()

	file, err = os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	writeNumbers(num, file)
}

func writeNumbers(num int, file *os.File) {
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println(num)
	file.WriteString(strconv.Itoa(num) + "\n")
}
