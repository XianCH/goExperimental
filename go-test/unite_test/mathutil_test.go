package unitetest

import "testing"

func TestFactoral(t *testing.T) {
	if result := Factorial(0); result != 1 {
		t.Errorf("Factorial is incorrect Got %d", result)
	}

	// 测试正整数的阶乘
	resutl := Factorial(5)
	expect := 120
	if resutl != expect {
		t.Errorf("Factorial is incorrect Got %d", resutl)
	}
}
