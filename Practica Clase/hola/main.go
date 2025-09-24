package main

import "fmt"

func suma(x int, y int) int {
	return x + y
}

func main() {
	var saludo string = "jojojo"
	saludo2 := "hola mundo" //Con los dos puntos antes del iguar ya asigna que es un string automaticamente.
	fmt.Println(saludo)
	fmt.Println(saludo2)

	x, y, z := 3, 4, 0

	z = x + y

	fmt.Printf("%d + %d = %d\n", x, y, z)

	array := [5]int{1, 2, 3, 4, 5}

	slc := array[1:3]

	slc = append(slc, 6)

	fmt.Println(slc)

	for i := 1; i <= 4; i++ {
		fmt.Println("poya y webos")
	}

	for pos, valor := range array {
		fmt.Printf("Posicion %d: %d\n", pos, valor)
	}

	continuar := true
	for continuar {
		fmt.Println("Hola mundo", continuar)
		continuar = false
	}

	var m map[int]string
	m = make(map[int]string)
	m[1] = "uno"
	m[2] = "dos"
	m[3] = "tres"
	m[4] = "cuatro"
	m[5] = "cinco"

	for clave, valor := range array {
		fmt.Printf("Posicion %d: %d\n", clave, valor)
	}

}
