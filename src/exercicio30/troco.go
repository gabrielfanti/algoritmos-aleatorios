package exercicio30

import (
	"fmt"
)

func Troco() {
	var total, paid float64

	fmt.Println("Cálculo de Troco")
	fmt.Print("Valor total a ser pago: R$")
	fmt.Scanln(&total)
	fmt.Print("Valor efetivamente pago: R$")
	fmt.Scanln(&paid)

	change := paid - total

	currencyUnits := map[string]float64{
		"R$100.00": 100.00,
		"R$50.00":  50.00,
		"R$10.00":  10.00,
		"R$5.00":   5.00,
		"R$1.00":   1.00,
		"R$0.50":   0.50,
		"R$0.10":   0.10,
		"R$0.05":   0.05,
		"R$0.01":   0.01,
	}

	fmt.Println("Troco: R$", change)

	for unit, value := range currencyUnits {
		if change >= value {
			count := int(change / value)
			change -= float64(count) * value
			fmt.Printf("%s: %d\n", unit, count)
		}
	}
}
