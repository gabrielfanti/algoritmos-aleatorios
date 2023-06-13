package exercicio24

import "fmt"

func JokenpoSheldon() string {
	var player1, player2 string

	fmt.Println("Jogo Jokenpo (Sheldon Cooper)")
	fmt.Println("Opções: Pedra, Papel, Tesoura, Lagarto, Spock")
	fmt.Println("Digite a jogada do Jogador 1 (Sheldon):")
	fmt.Scan(&player1)
	fmt.Println("Digite a jogada do Jogador 2:")
	fmt.Scan(&player2)

	if player1 == player2 {
		return "Empate"
	} else if (player1 == "Tesoura" && player2 == "Papel") ||
		(player1 == "Papel" && player2 == "Pedra") ||
		(player1 == "Pedra" && player2 == "Lagarto") ||
		(player1 == "Lagarto" && player2 == "Spock") ||
		(player1 == "Spock" && player2 == "Tesoura") ||
		(player1 == "Tesoura" && player2 == "Lagarto") ||
		(player1 == "Lagarto" && player2 == "Papel") ||
		(player1 == "Papel" && player2 == "Spock") ||
		(player1 == "Spock" && player2 == "Pedra") ||
		(player1 == "Pedra" && player2 == "Tesoura") {
		return "Jogador 1 (Sheldon) venceu"
	} else {
		return "Jogador 2 venceu"
	}
}
