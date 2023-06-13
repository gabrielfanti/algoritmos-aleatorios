package exercicio23

import "fmt"

func Jokenpo() string {
	var player1, player2 string

	fmt.Println("Jogo Jokenpo")
	fmt.Println("Opções: Pedra, Papel, Tesoura")
	fmt.Println("Digite a jogada do Jogador 1:")
	fmt.Scan(&player1)
	fmt.Println("Digite a jogada do Jogador 2:")
	fmt.Scan(&player2)

	if player1 == player2 {
		return "Empate"
	} else if (player1 == "Pedra" && player2 == "Tesoura") ||
		(player1 == "Tesoura" && player2 == "Papel") ||
		(player1 == "Papel" && player2 == "Pedra") {
		return "Jogador 1 venceu"
	} else {
		return "Jogador 2 venceu"
	}
}
