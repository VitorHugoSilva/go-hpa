package main

import (
	"testing"
)

func TestCalc(t *testing.T)  {
	caseTestes := []struct {
		valueTest float64
		result float64
	} {
		{.2, 0.6472135954999579},
		{.5, 1.2071067811865475},
		{2, 3.414213562373095},
	}
	for _, test := range caseTestes {
		t.Logf("Testando valor: %v \n",test.valueTest)
		caseTest := Calc(test.valueTest)

		if caseTest != test.result {
			t.Errorf("O valor esperadao é: %v, mas o retornado é %v", test.result, caseTest)
		}
	}
}