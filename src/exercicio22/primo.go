package exercicio22

import (
	"fmt"
	"math"
)

func Primo() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Printf("%d não é um número primo.\n", num)
		return
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			fmt.Printf("%d não é um número primo.\n", num)
			return
		}
	}

	fmt.Printf("%d é um número primo.\n", num)
}
