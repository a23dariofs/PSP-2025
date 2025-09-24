package main

import (
	"fmt"
	"sort"
)

// Tarea 1: Agrupar anagramas

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func agruparAnagrama(palabras []string) [][]string {
	agrupar := make(map[string][]string)

	for _, palabra := range palabras {
		palabraOrdenada := sortString(palabra)
		agrupar[palabraOrdenada] = append(agrupar[palabraOrdenada], palabra)
	}
	var resultado [][]string
	for _, grupo := range agrupar {
		resultado = append(resultado, grupo)
	}
	return resultado
}

func main() {
	palabras := []string{"hola", "gato", "peso", "alho", "toga"}

	grupos := agruparAnagrama(palabras)

	for _, grupo := range grupos {
		fmt.Println(grupo)
	}
}
