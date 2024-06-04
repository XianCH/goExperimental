package main

import (
	"fmt"

	"github.com/x14n/goExperimental/viperExper/config"
)

func main() {
	c := config.GetConfig()
	b := c.MySQL.Name
	fmt.Println(b)
	fmt.Println(c)
}
