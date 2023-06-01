package exercicio07

import (
	"fmt"
)

func SalarioLiquidoProfessor() {
	var valorHoraAula float64
	var horasTrabalhadas float64
	var percentualINSS float64

	fmt.Print("Informe o número de horas trabalhadas no mês: ")
	fmt.Scanln(&horasTrabalhadas)

	fmt.Print("Informe o valor da hora-aula: ")
	fmt.Scanln(&valorHoraAula)

	fmt.Print("Informe o percentual de desconto do INSS: ")
	fmt.Scanln(&percentualINSS)

	salarioBruto := valorHoraAula * horasTrabalhadas

	totalDesconto := salarioBruto * (percentualINSS / 100)

	salarioLiquido := salarioBruto - totalDesconto

	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}
