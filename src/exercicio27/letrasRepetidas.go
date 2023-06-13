package exercicio27

import (
	"fmt"
)

func LetrasRepetidas() {
	var text string

	fmt.Println("Encontrar a primeira letra não repetida")
	fmt.Print("Digite o texto: ")
	fmt.Scanln(&text)

	countMap := make(map[rune]int)

	for _, char := range text {
		countMap[char]++
	}

	for _, char := range text {
		if countMap[char] == 1 {
			fmt.Printf("A primeira letra não repetida no texto é: %s\n", string(char))
			return
		}
	}

	fmt.Println("Não existem letras que não se repetem no texto informado.")
}
