package exercicio31

import (
	"fmt"
	"strings"
)

func QuebraLinha() {
	var frase string
	var colunas int

	fmt.Println("Quebra de Linhas")
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&frase)
	fmt.Print("Digite a quantidade de colunas permitidas: ")
	fmt.Scanln(&colunas)

	palavras := strings.Split(frase, " ")
	linhas := []string{""}
	linhaAtual := 0

	for _, palavra := range palavras {
		if len(linhas[linhaAtual])+len(palavra) <= colunas {
			linhas[linhaAtual] += palavra + " "
		} else {
			linhas = append(linhas, palavra+" ")
			linhaAtual++
		}
	}

	fmt.Println("Resultado:")
	fmt.Println(strings.Join(linhas, "\n"))
}
