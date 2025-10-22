package main

import (
	"log"
	"os"
	"os/exec"
)

func StartAll(cmdList []*exec.Cmd) ([]*exec.Cmd, error) {

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
		cmd := exec.Command("go", "run", "child.go", string(rune('0'+i)))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
