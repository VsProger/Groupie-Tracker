package main

import (
	"fmt"
	"groupie-tracker/handlers"
)

func main() {
	fmt.Println("Server available on http://localhost:8080/")
	handlers.StartServer()
}
