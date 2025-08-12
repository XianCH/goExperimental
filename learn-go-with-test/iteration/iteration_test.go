package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkReapet(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

// func Benchmark(b *testing.B) {
// 	//... setup ...
// 	for b.Loop() {
// 		//... code to measure ...
// 	}
// 	//... cleanup ...
// }

func BenchmarkRepeatExample(b *testing.B) {
	for b.Loop() {
		RepeatExample("a")
	}
}

func BenchmarkRepeatExample02(b *testing.B) {
	for b.Loop() {
		RepeatExample02("a")
	}
}
