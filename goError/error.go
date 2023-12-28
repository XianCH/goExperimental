package main

import (
	"fmt"
)

type CustemError struct {
	Code    int
	Message string
	Data    any
}

func (e *CustemError) Error() string {
	return fmt.Sprintf("Error %d %s", e.Code, e.Message)
}

func SomeFunction() error {

	return &CustemError{
		Code:    500,
		Message: "Internal Server Error",
		Data:    map[string]interface{}{"key": "value"},
	}
}

func main() {
	err := SomeFunction()
	// 处理错误
	if err != nil {
		switch v := err.(type) {
		case *CustemError:
			fmt.Printf("Custom Error: Code=%d, Message=%s, Data=%v\n", v.Code, v.Message, v.Data)
		default:
			fmt.Println("Unknown Error:", err)
		}
	}
}
