package exercicio19

import (
	"fmt"
	"strings"
)

var tabelaPontuacao = map[rune]int{
	'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2, 'H': 4, 'I': 1, 'J': 8, 'K': 5, 'L': 1, 'M': 3,
	'N': 1, 'O': 1, 'P': 3, 'Q': 10, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 4, 'W': 4, 'X': 8, 'Y': 4, 'Z': 10,
}

func ScrabbleScore() {
	var word string

	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&word)

	word = strings.ReplaceAll(word, " ", "")
	word = strings.ToUpper(word)

	score := calcularScore(word)

	fmt.Printf("Score da palavra '%s': %d\n", word, score)
}

func calcularScore(palavra string) int {
	score := 0

	for _, letra := range palavra {
		score += tabelaPontuacao[letra]
	}

	return score
}
