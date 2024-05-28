package main

import (
	"github.com/x14n/goExperimental/viperExper/config"
)

func main() {
	config.TestViper()
	// _, err := config.InitRedis()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
