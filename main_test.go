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
	scenarios := []struct {
		input    string
		expected string
	}{
		{"10000\n500\n5\n10\n", "Após 10 anos, o valor do investimento será de R$ 94.111,23"},
		{"20000\n1000\n5\n10\n", "Após 10 anos, o valor do investimento será de R$ 188.222,47"},
		// Adicione mais cenários conforme necessário
	}

	for _, scenario := range scenarios {
		cmd := exec.Command("go", "run", ".")
		cmd.Stdin = strings.NewReader(scenario.input)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Falha ao executar o comando: %s", err)
		}
		output := strings.TrimSpace(string(out)) // Remove espaços em branco no início e no fim
		if !strings.Contains(output, scenario.expected) {
			t.Errorf("Teste falhou para entrada %s. Esperado '%s', Obtido '%s'", scenario.input, scenario.expected, output)
		}
	}
}
