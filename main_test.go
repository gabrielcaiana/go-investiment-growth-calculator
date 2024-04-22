package main

import (
	"testing"
)

// Testa a função calculaInvestimento
func TestCalculateInvestment(t *testing.T) {
	// Testa um cenário com valores conhecidos
	initial := 10000.0   // R$ 10.000 inicial
	monthly := 500.0     // R$ 500 por mês
	profitability := 5.0 // 5% ao ano
	years := 10          // por 10 anos

	expectValue := 94111.23

	// Chama a função que estamos testando
	result := calculateInvestment(initial, monthly, profitability, float64(years))

	// Verifica se o resultado está dentro de uma margem de erro aceitável
	if result < expectValue-1 || result > expectValue+1 {
		t.Errorf("Resultado incorreto: esperado ~%.2f, obtido %.2f", expectValue, result)
	}
}
