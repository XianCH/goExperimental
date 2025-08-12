package iteration

import "strings"

func Repeat(character string) string {
	result := make([]byte, 0, len(character)*5)

	for i := 0; i < 5; i++ {
		result = append(result, character...)
	}
	return string(result)
}

func RepeatExample(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

const repeatCount = 5

func RepeatExample02(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
