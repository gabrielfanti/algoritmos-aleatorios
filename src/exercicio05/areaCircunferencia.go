package exercicio05

import (
	"fmt"
	"math"
)

func AreaCircunferencia() {
	var raio float64
	const pi = 3.14159265

	fmt.Print("Digite o valor do raio: ")
	fmt.Scan(&raio)

	area := math.Pi * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é:", area)
}
