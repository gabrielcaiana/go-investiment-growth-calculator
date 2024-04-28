package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCalculateInvestment(t *testing.T) {
	tests := []struct {
		initial       float64
		monthly       float64
		profitability float64
		years         float64
		expected      float64
	}{
		{10000.0, 500.0, 5.0, 10, calculateInvestment(10000.0, 500.0, 5.0, 10)},
		{20000.0, 1000.0, 5.0, 10, calculateInvestment(20000.0, 1000.0, 5.0, 10)},
		{10000.0, 500.0, 10.0, 10, calculateInvestment(10000.0, 500.0, 10.0, 10)},
		{10000.0, 500.0, 5.0, 20, calculateInvestment(10000.0, 500.0, 5.0, 20)},
		{0.0, 500.0, 5.0, 10, calculateInvestment(0.0, 500.0, 5.0, 10)},
		{10000.0, 0.0, 5.0, 10, calculateInvestment(10000.0, 0.0, 5.0, 10)},
		{10000.0, 500.0, 0.0, 10, calculateInvestment(10000.0, 500.0, 0.0, 10)},
		{10000.0, 500.0, 5.0, 0, calculateInvestment(10000.0, 500.0, 5.0, 0)},
	}

	for _, test := range tests {
		result := calculateInvestment(test.initial, test.monthly, test.profitability, test.years)
		if result < test.expected-1 || result > test.expected+1 {
			t.Errorf("Para initial=%.2f, monthly=%.2f, profitability=%.2f, years=%.2f: esperado ~%.2f, obtido %.2f",
				test.initial, test.monthly, test.profitability, test.years, test.expected, result)
		}
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
