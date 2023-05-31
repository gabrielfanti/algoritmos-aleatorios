package exercicio14

import "fmt"

func EleicaoSindical() {
	var votosValidosA, votosValidosB, votosValidosC, votosNulos, votosEmBranco int

	fmt.Print("Digite a quantidade de votos válidos para o candidato A: ")
	fmt.Scanln(&votosValidosA)
	fmt.Print("Digite a quantidade de votos válidos para o candidato B: ")
	fmt.Scanln(&votosValidosB)
	fmt.Print("Digite a quantidade de votos válidos para o candidato C: ")
	fmt.Scanln(&votosValidosC)

	fmt.Print("Digite a quantidade de votos nulos: ")
	fmt.Scanln(&votosNulos)
	fmt.Print("Digite a quantidade de votos em branco: ")
	fmt.Scanln(&votosEmBranco)

	totalEleitores := votosValidosA + votosValidosB + votosValidosC + votosNulos + votosEmBranco

	percentVotosValidos := float64(votosValidosA+votosValidosB+votosValidosC) / float64(totalEleitores) * 100.0
	percentVotosA := float64(votosValidosA) / float64(totalEleitores) * 100.0
	percentVotosB := float64(votosValidosB) / float64(totalEleitores) * 100.0
	percentVotosC := float64(votosValidosC) / float64(totalEleitores) * 100.0
	percentVotosNulos := float64(votosNulos) / float64(totalEleitores) * 100.0
	percentVotosEmBranco := float64(votosEmBranco) / float64(totalEleitores) * 100.0

	fmt.Println("Resultado da votação:")
	fmt.Printf("Total de eleitores: %d\n", totalEleitores)
	fmt.Printf("Percentual de votos válidos: %.2f%%\n", percentVotosValidos)
	fmt.Printf("Percentual de votos válidos para o candidato A: %.2f%%\n", percentVotosA)
	fmt.Printf("Percentual de votos válidos para o candidato B: %.2f%%\n", percentVotosB)
	fmt.Printf("Percentual de votos válidos para o candidato C: %.2f%%\n", percentVotosC)
	fmt.Printf("Percentual de votos nulos: %.2f%%\n", percentVotosNulos)
	fmt.Printf("Percentual de votos em branco: %.2f%%\n", percentVotosEmBranco)
}
