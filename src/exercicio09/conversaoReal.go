package exercicio09

import (
	"fmt"
)

func ConversaoReal() {
	var cotacao float64
	var dolares float64

	fmt.Print("Digite a cotação do dólar: ")
	fmt.Scanln(&cotacao)

	fmt.Print("Digite a quantidade de dólares disponível: ")
	fmt.Scanln(&dolares)

	conversao := cotacao * dolares

	fmt.Printf("O valor em reais (R$) é: %.2f\n", conversao)
}
