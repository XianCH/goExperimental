package main

type topFunc func(v any) bool

func isEven(v any) bool {
	if num, ok := v.(int); ok {
		return num%2 == 0
	}
	return false
}

func isString(v any) bool {
	_, ok := v.(string)
	return ok
}

func isStringLenBigThanFive(v any) bool {
	if str, ok := v.(string); ok {
		return len(str) >= 5
	}
	return false
}

func filter(datas []any, f topFunc) []any {
	result := []any{}
	for _, data := range datas {
		if f(data) {
			result = append(result, data)
		}
	}
	return result
}

//
// func main() {
// 	datas := []any{1, 2, 4, 5, 6, "hello", "world", "foo"}
// 	result := filter(datas, isEven)
// 	fmt.Println(result)
//
// 	result = filter(datas, isString)
// 	fmt.Println(result)
// 	result = filter(datas, isStringLenBigThanFive)
// 	fmt.Println(result)
// }
