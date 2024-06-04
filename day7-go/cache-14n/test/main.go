package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

func main() {
	var sum int
	name := "x14n"

	for _, char := range name {
		sum += int(char)
	}

	fmt.Println(sum)
	fmt.Println(sum % 10)

	for i := 0; i < 3; i++ {
		fmt.Println(strconv.Itoa(i) + "key")
		fmt.Println(string(i) + "keys")
	}

	data := []byte("1keys")
	hash := int(crc32.ChecksumIEEE(data))
	fmt.Println(hash)
}
