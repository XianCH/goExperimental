package arraysandslices

// go test -cover

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		number := []int{1, 2, 3}
		got := Sum(number)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, number)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of two arrays", func(t *testing.T) {
		arrayA := [2]int{1, 2}
		arrayB := [2]int{3, 4}
		got := SumAll(arrayA, arrayB)
		want := [2]int{3, 7}
		// want := "boob"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given %v and %v", got, want, arrayA, arrayB)
		}
	})
}

func TestSumAllExmpel(t *testing.T) {
	t.Run("sum of two arrays", func(t *testing.T) {
		arrayA := []int{1, 2}
		arrayB := []int{3, 4}
		arrayC := []int{7, 7}
		got := SumAllExmple(arrayA, arrayB, arrayC)
		want := []int{3, 7, 14}
		// want := "boob"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given %v and %v", got, want, arrayA, arrayB)
		}
	})
}
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(numbers)
	}
}
