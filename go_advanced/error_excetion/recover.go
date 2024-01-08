package main

import (
	"log"
)

func test_recover() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
}
