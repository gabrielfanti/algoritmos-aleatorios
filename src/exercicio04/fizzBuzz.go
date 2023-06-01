package exercicio04

import (
	"fmt"
	"strconv"
)

func FizzBuzz() {
	var start, end int
	fmt.Print("Informe o número inicial: ")
	fmt.Scan(&start)
	fmt.Print("Informe o número final: ")
	fmt.Scan(&end)

	if end <= start {
		fmt.Println("O número final deve ser maior que o número inicial.")
		return
	}

	for i := start; i <= end; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		fmt.Println(output)
	}
}
