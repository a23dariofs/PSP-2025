package main

import (
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"strconv"
)

func StartAll(cmdList []*exec.Cmd) ([]*exec.Cmd, error) {
	for range 20 {
		i, j := rand.IntN(10), rand.IntN(10)
		cmdList[i], cmdList[j] = cmdList[j], cmdList[i]
	}
	for _, cmd := range cmdList {
		err := cmd.Start()
		if err != nil {
			return nil, err
		}
	}
	return cmdList, nil
}

func main() {
	var cmdList []*exec.Cmd

	for i := 0; i < 10; i++ {
		cmd := exec.Command("go", "run", "hijo/child.go", strconv.Itoa(i))
		cmdList = append(cmdList, cmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	cmdList, err := StartAll(cmdList)
	if err != nil {
		log.Fatal("Something went wrong:", err)
	}

	for _, cmd := range cmdList {
		cmd.Wait()
	}

}
