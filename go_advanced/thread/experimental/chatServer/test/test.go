package main

import (
	"fmt"
	"strings"
	"time"
)

func testByte() {
	time := time.Now().Format("15:04:05")
	username := "xianchaohao"
	str := time + "/" + "hello" + "/" + username
	bytedata := []byte(str)
	fmt.Println(string(bytedata))
	strSlice := strings.Split(str, "/")
	fmt.Println(strSlice)

	for _, value := range strSlice {
		fmt.Println(value)
	}
}

func main() {
	testByte()
}
