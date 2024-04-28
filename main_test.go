package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCalculateInvestment(t *testing.T) {
	initial := 10000.0   // R$ 10.000 inicial
	monthly := 500.0     // R$ 500 por mês
	profitability := 5.0 // 5% ao ano
	years := 10          // por 10 anos

	expectValue := 94111.23

	result := calculateInvestment(initial, monthly, profitability, float64(years))

	if result < expectValue-1 || result > expectValue+1 {
		t.Errorf("Resultado incorreto: esperado ~%.2f, obtido %.2f", expectValue, result)
	}
}

func TestMain(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader("1000\n200\n0.9\n5\n")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Expected no error, but got '%s'", err)
	}

	expected := "Após 5 anos, o valor do investimento será de R$ 13.315,40\n"
	if !strings.Contains(string(out), expected) {
		t.Errorf("Expected '%s' but got '%s'", expected, string(out))
	}
}
