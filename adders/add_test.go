package integers

import "testing"

func TestAdder (t *testing.T) {
	sum := Add(3,2)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d', got '%d'", expected, sum)
	}
}