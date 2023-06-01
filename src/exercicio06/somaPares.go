package exercicio06

import (
	"fmt"
)

func SomaPares() {
	var inicio, fim int
	fmt.Println("Digite o número inicial do intervalo:")
	fmt.Scanln(&inicio)
	fmt.Println("Digite o número final do intervalo:")
	fmt.Scanln(&fim)

	soma := 0
	for i := inicio; i <= fim; i++ {
		if i%2 == 0 {
			soma += i
		}
	}

	fmt.Println("A soma dos números pares no intervalo é:", soma)
}
