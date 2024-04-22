package main

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Calcula o valor futuro de um investimento com depósitos regulares
func calculateInvestment(initial float64, monthly float64, profitability float64, years float64) float64 {
	// Convertendo a taxa anual para mensal
	taxMonthly := profitability / 12 / 100

	// Calcula o valor total de meses
	totalMonthly := years * 12

	// Valor bruto do investimento
	futureValue := initial * math.Pow(1+taxMonthly, float64(totalMonthly))

	for i := 0; i < int(totalMonthly); i++ {
		futureValue += monthly * math.Pow(1+taxMonthly, totalMonthly-float64(i)-1)
	}

	return futureValue

}

func main() {
	// Declaração das variáveis que armazenarão as entradas do usuário
	var initial float64
	var monthly float64
	var profitability float64
	var years float64

	// Coleta os dados do usuário
	fmt.Print("Informe o valor do investimento\n")
	fmt.Scan(&initial)
	fmt.Print("Informe o valor do depósito mensal\n")
	fmt.Scan(&monthly)
	fmt.Print("Informe a taxa de rentabilidade anual\n")
	fmt.Scan(&profitability)
	fmt.Print("Informe o período de investimento em anos\n")
	fmt.Scan(&years)

	// Calculando a rentabilidade do investimento
	result := calculateInvestment(initial, monthly, profitability, years)

	// Criando um novo Printer para o idioma português brasileiro
	p := message.NewPrinter(language.BrazilianPortuguese)

	// Exibindo o resultado
	fmt.Printf("Após %d anos, o valor do investimento será de ", int(years))
	p.Printf("R$ %.2f\n", result)
}
