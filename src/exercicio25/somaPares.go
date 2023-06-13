package exercicio25

import (
	"fmt"
)

func SomaPares() {
	var numbers []float64

	fmt.Println("Digite os números desejados (ou digite 'sair' para encerrar):")
	for {
		var input string
		fmt.Scan(&input)

		if input == "sair" {
			break
		}

		var num float64
		_, err := fmt.Sscan(input, &num)
		if err != nil {
			fmt.Println("Entrada inválida. Digite um número válido ou 'sair' para encerrar.")
			continue
		}

		numbers = append(numbers, num)
	}

	if len(numbers) > 0 {
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		average := sum / float64(len(numbers))
		fmt.Printf("A média dos números informados é: %.2f\n", average)
	} else {
		fmt.Println("Nenhum número foi informado.")
	}
}
