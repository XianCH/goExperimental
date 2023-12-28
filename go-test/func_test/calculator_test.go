package functest

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("Add result is incorrent.Got %d", ans)
	}
}

func TestSub(t *testing.T) {
	if ans := Sub(3, 1); ans != 2 {
		t.Errorf("Sub result is incorrent.Got %d", ans)
	}
}
