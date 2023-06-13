package exercicio21

import "fmt"

func AlgarismosRomanos() int {
	numeral := "MCMXCIV"

	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(numeral)

	for i := 0; i < n; i++ {
		if i+1 < n && romanValues[numeral[i]] < romanValues[numeral[i+1]] {
			result -= romanValues[numeral[i]]
		} else {
			result += romanValues[numeral[i]]
		}
	}

	fmt.Printf("O numeral romano %s corresponde ao número inteiro %d\n", numeral, result)
	return result
}
