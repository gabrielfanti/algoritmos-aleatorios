package exercicio03

import (
	"fmt"
)

func AjusteSalarial() {
	var salario float64
	var percentual float64

	fmt.Print("Digite o salário mensal: ")
	fmt.Scan(&salario)

	fmt.Print("Digite o percentual de reajuste: ")
	fmt.Scan(&percentual)

	novoSalario := salario + (salario * percentual / 100)

	fmt.Printf("O novo salário é: R$ %.2f\n", novoSalario)
}
