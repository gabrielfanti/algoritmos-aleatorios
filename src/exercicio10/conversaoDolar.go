package exercicio10

import "fmt"

func ConversaoDolar() {
	var cotacao float64
	var valorReal float64

	fmt.Print("Informe a cotação do dólar: ")
	fmt.Scan(&cotacao)

	fmt.Print("Informe a quantidade de reais disponível: ")
	fmt.Scan(&valorReal)

	valorDolar := valorReal / cotacao

	fmt.Printf("O valor em dólar é: US$ %.2f\n", valorDolar)
}
