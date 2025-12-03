package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

func main() {
	palabras := []string{"kayak", "pito", "noon", "peep", "juan", "psp", "lamineyamal", "huevitos"}

	var wg sync.WaitGroup
	for i := 0; i < len(palabras); i++ {
		wg.Add(1)
		go isPalindrome(palabras[i], &wg)
	}
	wg.Wait()
}

func isPalindrome(word string, wg *sync.WaitGroup) {
	defer wg.Done()

	var letras = strings.Split(word, "")
	slices.Reverse(letras)
	var palindrome = strings.Join(letras, "")

	if palindrome == word {
		fmt.Println(word, "es Palindrome")
	} else {
		fmt.Println(word, "no es Palindrome")
	}

}
