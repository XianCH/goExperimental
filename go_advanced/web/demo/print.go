package main

import "fmt"

type Printer interface {
	Print(message string)
}

type PrintFunc func(message string)

func (p PrintFunc) Print(message string) {
	p(message)
}

func Greet(p Printer, msg string) {
	p.Print(msg)
}

func main() {
	printHello := PrintFunc(func(message string) {
		fmt.Println("hello" + message)
	})
	Greet(printHello, "World")
}
