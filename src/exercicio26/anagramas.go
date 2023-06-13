package exercicio26

import (
	"fmt"
	"sort"
	"strings"
)

func Anagramas() {
	var str1, str2 string

	fmt.Println("Verificação de Anagramas")
	fmt.Print("Digite a primeira string: ")
	fmt.Scan(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scan(&str2)

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	str1Chars := strings.Split(str1, "")
	str2Chars := strings.Split(str2, "")

	sort.Strings(str1Chars)
	sort.Strings(str2Chars)

	str1Sorted := strings.Join(str1Chars, "")
	str2Sorted := strings.Join(str2Chars, "")

	if str1Sorted == str2Sorted {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
