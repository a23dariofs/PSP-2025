# COSAS A SABER DEL EJERCICIO:

os → para manejar archivos y argumentos de línea de comando.

#### if len(os.Args) < 2 { return }

Revisa que se haya pasado un argumento al programa (el número).

os.Args[0] es el nombre del programa, os.Args[1] será el número que pasa el padre.

#### numStr := os.Args[1]

Guarda el número pasado por el padre como string.

#### fmt.Println(numStr)

Imprime en consola una sola línea con el número.

#### file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 

1. os.O_APPEND → agrega al final sin borrar el contenido.

2. os.O_CREATE → crea el archivo si no existe.

3. os.O_WRONLY → abre el archivo solo para escribir.

4. 0644 → permisos de lectura/escritura.

#### defer file.Close()

Asegura que el archivo se cierre al terminar la función.

#### file.WriteString(numStr + "\n")

Escribe el número en el archivo, seguido de un salto de línea.


´´´

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



package main

import (
	"fmt"
	"os"
)

func main() {
	//Lee que se hayan pasado todos los argumentos y si no termina el programa (son dos porque el args[0] es el nombre del programa y args[1] el argumento en string)
	if len(os.Args) < 2 {
		return
	}

	//Añadimos a una variable llamada numStr el argumento pasado en el padre y lo imprimimos
	numStr := os.Args[1]

	fmt.Println(numStr)

	//Creamos el file
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	//Cierra el file
	defer file.Close()

	//Esto es para escribir en el file cada uno de los numeros con un salto de linea
	file.WriteString(numStr + "\n")
}