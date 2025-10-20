package main

import (
	"fmt"
	"os/exec"
)

//Ejecutar comandos en cmd

func main() {
	cmd := exec.Command("cmd", "/C", "dir")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		return
	}
	fmt.Println(string(output))
}

// abrir una aplicacion (Bloc de notas)

/*
func main() {
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error al abrir Notepad:", err)
		return
	}
	fmt.Println("Notepad se ha abierto como un nuevo proceso.")
}

*/

// Ejecutar varios procesos externos en paralelo

/*
func main() {
	comandos := []string{"echo Hola", "time /T", "whoami"}
	var wg sync.WaitGroup

	for _, cmdStr := range comandos {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			cmd := exec.Command("cmd", "/C", c)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error ejecutando", c, ":", err)
				return
			}
			fmt.Printf("Salida de '%s':\n%s\n", c, string(out))
		}(cmdStr)
	}

	wg.Wait()
	fmt.Println("Todos los comandos han terminado.")
}

*/
