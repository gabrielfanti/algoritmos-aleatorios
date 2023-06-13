package exercicio28

import (
	"fmt"
)

func Fibonacci() {
	var num int

	fmt.Println("Sequência de Fibonacci")
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	sequence := []int{0, 1}

	for i := 2; sequence[i-1]+sequence[i-2] <= num; i++ {
		nextNum := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, nextNum)
	}

	fmt.Println("Sequência de Fibonacci até", num, ":")
	for _, num := range sequence {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
