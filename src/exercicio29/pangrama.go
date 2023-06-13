package exercicio29

import (
	"fmt"
	"strings"
)

func Pangrama() {
	letras := make(map[rune]bool)
	var frase string
	var validacao bool

	fmt.Println("Verificação de Pangrama")
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&frase)

	frase = strings.ReplaceAll(frase, " ", "")
	frase = strings.ToLower(frase)

	for _, char := range frase {
		letras[char] = true
	}
	validacao = true
	for char := 'a'; char <= 'z'; char++ {
		if !letras[char] {
			validacao = false
			break
		}
	}

	if validacao {
		fmt.Println("A frase é um pangrama.")
	} else {
		fmt.Println("A frase não é um pangrama.")
	}
}
