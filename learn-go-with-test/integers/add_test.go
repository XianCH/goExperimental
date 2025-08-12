package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("expected '%d' but got '%d'", expect, sum)
	}
}
