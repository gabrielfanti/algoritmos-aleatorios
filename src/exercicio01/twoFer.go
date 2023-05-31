package exercicio01

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TwoFer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um nome: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		name = "you"
	}

	result := fmt.Sprintf("One for %s, one for me.", name)
	fmt.Println(result)
}
