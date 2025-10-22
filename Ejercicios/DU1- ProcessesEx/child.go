package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	numStr := os.Args[1]
	num, _ := strconv.Atoi(numStr)

	fmt.Println(num)

	var flag int
	if num == 0 {
		flag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	} else {
		flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	}

	file, err := os.OpenFile("output.txt", flag, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%d\n", num))
}
