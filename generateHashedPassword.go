package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := "12345"
	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(pw), 12)
	fmt.Println(string(hashedPw))
}
