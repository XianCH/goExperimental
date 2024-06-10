package main

import (
	"fmt"

	jwtexp "github.com/x14n/goExperimental/jwt/tokenExp/jwtExp"
)

func main() {

	s, err := jwtexp.GenToken("xianchaoaho")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)

	mc, err2 := jwtexp.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cI6IkpXVCJ9.J1c2VybmFtZSI6InhpYW5jaGFvYWhvIiwiZXhwIjoxNzAwNzQyMzg4LCJpc3MiOiJteS1wcm9qZWN0In0.99JooV-iF1265WYfmc72VTwm8V8MvdBqagtmfdvn8Ss")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(mc)
}
