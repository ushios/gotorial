package calc

import "testing"

func TestDivision(t *testing.T) {
	var a, b int
	a = 10
	b = 5

	div, err := Division(a, b)

	if err != nil {
		t.Errorf("Division has error.")
	}

	if div != 2 {
		t.Errorf("sum want (%d).", 2)
	}
}

func TestDivisionNoErr(t *testing.T) {
	var a, b int
	a = 20
	b = 5

	div, _ := Division(a, b)

	if div != 4 {
		t.Errorf("sum want (%d).", 4)
	}
}
