package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$T8Z.H4tLh3/K6U1eBx5/8Oof3XfMwB26MkV7kCWe8b.vI4l9H4IYi"
	
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("admin123"))
	fmt.Println("Does it match admin123?", err == nil, err)
	
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte("password"))
	fmt.Println("Does it match password?", err == nil, err)

	newHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	fmt.Println("New Hash for admin123:", string(newHash))
}
