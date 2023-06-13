package exercicio20

import "fmt"

func Fatorial() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	var calculo func(int) int
	calculo = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * calculo(n-1)
	}

	result := calculo(num)
	fmt.Printf("O fatorial de %d é %d\n", num, result)
}
