package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func StartAll(cmdList []*exec.Cmd) ([]*exec.Cmd, error) {
	// Lanzar los comandos en el orden que están en la lista
	for _, cmd := range cmdList {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			return nil, err
		}
	}
	return cmdList, nil
}

func main() {
	for i := 0; i < 10; i++ {
		cmd := exec.Command("go", "run", "child.go", fmt.Sprintf("%d", i))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
