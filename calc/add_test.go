package calc

import "testing"

func TestAdd(t *testing.T) {
	var a, b int
	a = 10
	b = 15

	sum := Add(a, b)

	if sum != 25 {
		t.Errorf("sum want (%d)", 25)
	}
}

func TestThreeAdd(t *testing.T) {
	var a, b, c int
	a = 10
	b = 20
	c = 30

	sum := threeAdd(a, b, c)

	if sum != 60 {
		t.Errorf("sum want (%d)", 60)
	}
}
