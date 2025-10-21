package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {

	//Reinicia el archivo para que estea vacio una vez ejecutemos el programa y asi la salida siempre será la pedida
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	file.Close()

	//Ejecuta 10 procesos hijo
	for i := 0; i < 10; i++ {
		runChild(i)
	}
}

func runChild(n int) {

	//Comando de la terminal que llama al hijo y le pasa el argumento en string ya que la terminal en Go solo acepta strings
	cmd := exec.Command("go", "run", "child.go", string('0'+n))

	//Captura la salida estándar del hijo para que el padre pueda leerla y si tiene errores los manda a la terminal
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = os.Stderr
	cmd.Start()

	//Copia la salida del hijo a la terminal del padre asi evitamos usar fmt.Print
	io.Copy(os.Stdout, stdout)

	//Espera a que el hijo termine antes de continuar
	cmd.Wait()
}
