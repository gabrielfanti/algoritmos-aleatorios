package exercicio02

import "fmt"

func Troca() {
	// Leitura dos valores de A e B
	var A, B int
	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&A)
	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&B)

	// Troca dos valores
	A, B = B, A

	// Apresentação dos valores após a troca
	fmt.Println("Valores após a troca:")
	fmt.Println("A =", A)
	fmt.Println("B =", B)
}
