package exercicio08

import "fmt"

func VolumeCaixa() {
	// Solicitar as dimensões da caixa ao usuário
	fmt.Println("Digite o comprimento da caixa:")
	var comprimento float64
	fmt.Scanln(&comprimento)

	fmt.Println("Digite a largura da caixa:")
	var largura float64
	fmt.Scanln(&largura)

	fmt.Println("Digite a altura da caixa:")
	var altura float64
	fmt.Scanln(&altura)

	// Calcular o volume
	volume := comprimento * largura * altura

	// Apresentar o resultado
	fmt.Println("O volume da caixa é:", volume)
}
