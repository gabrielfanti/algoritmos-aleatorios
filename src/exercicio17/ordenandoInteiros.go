package exercicio17

import (
	"fmt"
	"sort"
)

func OrdenandoInteiros() {
	var elementos [12]int

	for i := 0; i < 12; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scanln(&elementos[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elementos[:])))

	fmt.Println("Elementos em ordem decrescente:")
	for _, elemento := range elementos {
		fmt.Println(elemento)
	}
}
