package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/Chayathon/go-example/chayathon"
)

func generateUUID() string {
	id := uuid.New()
	return id.String()
}

func main () {
	hello := "Hello, World!"
	fmt.Println(hello)
	fmt.Println("UUID: ", generateUUID())
	chayathon.SayHello()
	chayathon.SayHelloGolang()
}