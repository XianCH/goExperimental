package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	newUUID := uuid.New()
	fmt.Println(newUUID)

	newUUIDsTRING := uuid.New().String()

	fmt.Println(newUUIDsTRING)
	
}
