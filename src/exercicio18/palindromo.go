package exercicio18

import (
	"fmt"
	"strings"
)

func Palindromo() {
	var text string

	fmt.Print("Digite uma palavra, frase ou número: ")
	fmt.Scanln(&text)

	text = strings.ReplaceAll(text, " ", "")

	if isPalindromo(text) {
		fmt.Println("A entrada é um palíndromo!")
	} else {
		fmt.Println("A entrada não é um palíndromo.")
	}
}

func isPalindromo(str string) bool {
	str = strings.ToLower(str)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
