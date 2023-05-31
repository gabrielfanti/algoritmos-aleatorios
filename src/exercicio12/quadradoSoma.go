package exercicio12

import (
	"fmt"
	"math"
)

func QuadradoSoma() {
	var num1, num2, num3 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	soma := num1 + num2 + num3
	resultado := int(math.Pow(float64(soma), 2))

	fmt.Printf("O resultado é: %d\n", resultado)
}
