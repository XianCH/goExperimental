package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrayA, arrayB [2]int) [2]int {
	var sum [2]int
	sum[0] = Sum(arrayA[:])
	sum[1] = Sum(arrayB[:])
	return sum
}

func SumAllExmple(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
