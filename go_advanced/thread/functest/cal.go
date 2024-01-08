package main

type cal func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func factor(a, b int) int {
	return a * b
}

func Calculate(a int, b int, calcal cal) int {
	return calcal(a, b)
}

//
// func main() {
// 	//add
// 	fmt.Println(Calculate(1, 2, add))
//
// 	//sub
// 	fmt.Println(Calculate(1, 2, sub))
//
// 	//factor
// 	fmt.Println(Calculate(2, 2, factor))
// }
